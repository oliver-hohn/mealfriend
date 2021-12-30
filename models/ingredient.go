package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Ingredient struct {
	gorm.Model
	Name string
	Type sql.NullString

	RecipeID uint64
	Recipe   Recipe
}
