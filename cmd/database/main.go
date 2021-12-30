package main

import (
	"fmt"
	"log"

	"github.com/oliver-hohn/mealfriend/database"
	"github.com/oliver-hohn/mealfriend/envs"
)

func main() {
	conn, err := database.CreateConn(database.DatabaseConfig{
		Host:     envs.MustGetEnv("PGHOST"),
		Port:     envs.MustGetIntEnv("PGPORT"),
		Database: envs.MustGetEnv("PGDATABASE"),
		Username: envs.MustGetEnv("PGUSER"),
		Password: envs.MustGetEnv("PGPASSWORD"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Exec("select 1").Error; err != nil {
		log.Fatalf("unable to query: %v", err)
	}

	fmt.Println("successful DB setup!")
}
