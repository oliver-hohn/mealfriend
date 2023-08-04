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

### Scrape a recipe
1. Run:
   ```sh
   bin/docker_run go run cmd/scrape/main.go --input_url=RECIPE_URL
   ```
  _To store the recipe, use the `--store=1` option_:
  ```sh
   bin/docker_run go run cmd/scrape/main.go --store=1 --input_url=RECIPE_URL
   ```