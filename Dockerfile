FROM golang:1.20 as base

FROM base as dev
  WORKDIR /app

  # Install air for hot-reloads in development: https://github.com/cosmtrek/air#installation
  RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

  # Go debugger
  RUN go install github.com/go-delve/delve/cmd/dlv@latest

  CMD ["air"]
