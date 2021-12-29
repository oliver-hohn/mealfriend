package main

import (
	"context"
	"fmt"
	"log"

	"github.com/oliver-hohn/mealfriend/database"
	"github.com/oliver-hohn/mealfriend/envs"
)

var DB_URL = envs.MustGetEnv("DATABASE_URL")

func main() {
	ctx := context.Background()

	conn, err := database.Start(DB_URL, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	if _, err := conn.Exec(ctx, "select 1"); err != nil {
		log.Fatalf("unable to query: %v", err)
	}

	fmt.Println("successful DB setup!")
}
