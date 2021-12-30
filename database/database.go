package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     int64
	Database string
	Username string
	Password string
}

func CreateConn(dbConfig DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.Username, dbConfig.Password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %w", err)
	}

	return db, nil
}
