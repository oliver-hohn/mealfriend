package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/url"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/oliver-hohn/mealfriend/envs"
	"github.com/oliver-hohn/mealfriend/graph"
	"github.com/oliver-hohn/mealfriend/models"
	pb "github.com/oliver-hohn/mealfriend/protos"
	"github.com/oliver-hohn/mealfriend/scrapers"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"

	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

type mealfriendServer struct {
	pb.UnimplementedMealfriendServer

	driver neo4j.DriverWithContext
}

func (s *mealfriendServer) Scrape(ctx context.Context, req *pb.ScrapeRequest) (*pb.ScrapeResponse, error) {
	if len(req.GetUrl()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing required argument: url")
	}

	u, err := url.Parse(req.GetUrl())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("cannot parse %s: %v", req.GetUrl(), err))
	}

	recipe, err := scrapers.Scrape(u)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("unable to scrape: %v", err))
	}

	if err := graph.SaveRecipe(ctx, s.driver, recipe); err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("unable to write %s, due to: %v", recipe.Code, err))
	}

	return &pb.ScrapeResponse{Recipe: toProto(recipe)}, nil
}

const MAX_PLAN_SIZE = 5

func (s *mealfriendServer) GetMealPlan(ctx context.Context, req *pb.GetMealPlanRequest) (*pb.GetMealPlanResponse, error) {
	// Validate that the number of requested recipes does not exceed the limit
	requested_recipes := 0
	for _, n := range req.GetRequirements() {
		requested_recipes += int(n)
	}
	if requested_recipes > MAX_PLAN_SIZE {
		return nil, status.Error(codes.FailedPrecondition, fmt.Sprintf("requested too many recipes (max: %d, requested: %d)", MAX_PLAN_SIZE, requested_recipes))
	}

	protoRecipes := []*pb.Recipe{}
	recipeCodes := []string{}
	for t, n := range req.GetRequirements() {
		recipes, err := graph.FindRecipes(ctx, s.driver, models.Tag(t), int(n), recipeCodes)
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("unable to find recipes for %s, due to: %v", t, err))
		}

		for _, r := range recipes {
			protoRecipes = append(protoRecipes, toProto(r))
			recipeCodes = append(recipeCodes, r.Code)
		}
	}

	return &pb.GetMealPlanResponse{Recipes: protoRecipes}, nil
}

func toProto(r *models.Recipe) *pb.Recipe {
	tags := make([]pb.Tag, len(r.Tags))

	for i, t := range r.Tags {
		tags[i] = toProtoTag(t)
	}

	return &pb.Recipe{
		Code:        r.Code,
		Name:        r.Name,
		Source:      r.Source.String(),
		Ingredients: r.Ingredients,
		Tags:        tags,
		CookTime:    &durationpb.Duration{Seconds: int64(r.CookTime.Seconds())},
	}
}

func toProtoTag(t models.Tag) pb.Tag {
	switch t {
	case models.BEEF:
		return pb.Tag_TAG_BEEF
	case models.DAIRY:
		return pb.Tag_TAG_DAIRY
	case models.EGG:
		return pb.Tag_TAG_EGG
	case models.FISH:
		return pb.Tag_TAG_FISH
	case models.FRUIT:
		return pb.Tag_TAG_FRUIT
	case models.GRAIN:
		return pb.Tag_TAG_GRAIN
	case models.LEGUMES:
		return pb.Tag_TAG_LEGUMES
	case models.PASTA:
		return pb.Tag_TAG_PASTA
	case models.PORK:
		return pb.Tag_TAG_PORK
	case models.POULTRY:
		return pb.Tag_TAG_POULTRY
	case models.SHELLFISH:
		return pb.Tag_TAG_SHELLFISH
	case models.VEGETABLE:
		return pb.Tag_TAG_VEGETABLE
	default:
		return pb.Tag_TAG_UNSPECIFIED
	}
}

func InterceptorLogger(l *zap.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make([]zap.Field, 0, len(fields)/2)

		for i := 0; i < len(fields); i += 2 {
			key := fields[i]
			value := fields[i+1]

			switch v := value.(type) {
			case string:
				f = append(f, zap.String(key.(string), v))
			case int:
				f = append(f, zap.Int(key.(string), v))
			case bool:
				f = append(f, zap.Bool(key.(string), v))
			default:
				f = append(f, zap.Any(key.(string), v))
			}
		}

		logger := l.WithOptions(zap.AddCallerSkip(1)).With(f...)

		switch lvl {
		case logging.LevelDebug:
			logger.Debug(msg)
		case logging.LevelInfo:
			logger.Info(msg)
		case logging.LevelWarn:
			logger.Warn(msg)
		case logging.LevelError:
			logger.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func main() {
	logger := zap.NewExample()

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(envs.MustGetEnv("NEO4J_URI"), neo4j.BasicAuth(envs.MustGetEnv("NEO4J_USERNAME"), envs.MustGetEnv("NEO4J_PASSWORD"), ""))
	if err != nil {
		log.Fatalf("unable to initialize neo4j driver: %v", err)
	}
	defer driver.Close(ctx)

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(InterceptorLogger(logger), opts...),
		),
	)
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(s, healthcheck)
	pb.RegisterMealfriendServer(s, &mealfriendServer{driver: driver})
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
