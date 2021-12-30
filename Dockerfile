FROM golang:1.17 as base

FROM base as dev
  WORKDIR /app

  # Install the Postgres Client to dump the DB schema
  RUN apt-get update -qq \
    && apt-get install -y --no-install-recommends postgresql-client-13 \
    && rm -rf /var/lib/apt/lists/*

  # Install Goose to manage database migrations
  RUN go install github.com/pressly/goose/v3/cmd/goose@latest
  
  # Install air for hot-reloads in development: https://github.com/cosmtrek/air#installation
  RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
  
  CMD ["air"]
