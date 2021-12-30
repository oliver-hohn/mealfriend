package main

import (
	"context"
	"fmt"
	"log"

	"github.com/oliver-hohn/mealfriend/database"
	"github.com/oliver-hohn/mealfriend/envs"
)

func main() {
	ctx := context.Background()

	conn, err := database.Start(mustGenerateDatabaseURL(), ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	if _, err := conn.Exec(ctx, "select 1"); err != nil {
		log.Fatalf("unable to query: %v", err)
	}

	fmt.Println("successful DB setup!")
}

func mustGenerateDatabaseURL() string {
	host := envs.MustGetEnv("PGHOST")
	port := envs.MustGetEnv("PGPORT")
	database := envs.MustGetEnv("PGDATABASE")
	user := envs.MustGetEnv("PGUSER")
	password := envs.MustGetEnv("PGPASSWORD")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
}
