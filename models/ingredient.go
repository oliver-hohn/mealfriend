package models

import (
	"gorm.io/gorm"
)

type Ingredient struct {
	gorm.Model
	Name string

	RecipeID uint64
	Recipe   Recipe
}
