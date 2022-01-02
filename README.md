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
1. Run:
   ```sh
   bin/docker_run go run cmd/plan/main.go --count=3
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

## Create DB migrations
1. Run:
  ```sh
  bin/docker_goose create name_of_migration sql
  ```
1. Run:
  ```
  bin/docker_goose up
  ```
