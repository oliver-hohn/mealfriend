# Mealfriend

## Setup
1. Install Docker and Docker-Compose: https://docs.docker.com/get-docker/.
1. Build the container:
   ```
   docker-compose build
   ```
1. Install `grpcurl`: https://github.com/fullstorydev/grpcurl
1. Run (to seed the local DB):
   ```sh
   bin/docker_run go run cmd/seed/main.go --seed_file=config/seed.csv
   ```

## Usage
1. Run:
   ```
   docker-compose up
   ```
1. In a separate terminal, run:
   ```
   cd frontend/

   npm install
   npm run dev
   ```
1. Go to: `http://localhost:3000`

### gRPC
Seed a recipe:
```sh
grpcurl -plaintext -d @ localhost:50051 mealfriend.Mealfriend/Scrape <<EOM
{
   "url": "https://cafedelites.com/best-churros-recipe/"
}
```

Plan meals:
```sh
grpcurl -plaintext -d @ localhost:50051 mealfriend.Mealfriend/GetMealPlan <<EOM
{
   "requirements": {
      "beef": 1,
      "poultry": 1,
      "fish": 2,
      "unspecified": 1
   }
}
EOM
```
_`unspecified` acts as a "filler" for any recipe (i.e. no requirement)._

### Scrape a recipe
1. Run:
   ```sh
   bin/docker_run go run cmd/scrape/main.go --input_url=RECIPE_URL
   ```
   _To store the recipe, use the `--store=1` option_:
   ```sh
   bin/docker_run go run cmd/scrape/main.go --store=1 --input_url=RECIPE_URL
   ```

## Visualising the DB graph locally (Neo4j)
1. Run: `docker-compose up`
1. Go to: http://localhost:7474/browser/
   1. Connect to: `bolt://localhost:7687`, grab the username and password from docker-compose.yml

## Re-generating the protos
1. Install `protobuf` and `protoc-gen-grpc-web`:
   ```
   brew install protobuf
   brew install protoc-gen-grpc-web
   npm i -g ts-protoc-gen
   ```
1. Run:
   ```sh
   protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      --js_out=import_style=commonjs,binary:./frontend/src \
      --ts_out=service=grpc-web:./frontend/src \
      protos/mealfriend.proto
   ```

## gRPC-web
1. Start the container:
   ```
   docker-compose up
   ```
1. In a separate console, run:
   ```
   cd frontend/
   npm run dev
   ```
1. Go to: `localhost:3000`, and open the console. You should see successful requests to the gRPC server.

## Debug
1. Run:
   ```sh
   bin/docker_run dlv debug PACKAGE_NAME -- -arg1=val1
   ```
   e.g.
   ```sh
   bin/docker_run dlv debug github.com/oliver-hohn/mealfriend/cmd/seed -- -seed_file=config/seed.csv
   ```
