# Mealfriend

## Setup
1. Install Docker and Docker-Compose: https://docs.docker.com/get-docker/.
1. Build the container:
   ```
   docker-compose build
   ```

## Usage
### Seed
1. Run:
   ```sh
   bin/docker_run go run cmd/seed/main.go --seed_file=config/seed.csv
   ```

### Plan
1. Run (after seeding):
   ```sh
   bin/docker_run go run cmd/plan/main.go --poultry=1 --fish=2
   ```

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