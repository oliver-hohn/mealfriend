package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type IngredientType string

const (
	BEEF      IngredientType = "BEEF"
	CHICKEN   IngredientType = "CHICKEN"
	PORK      IngredientType = "PORK"
	FISH      IngredientType = "FISH"
	VEGETABLE IngredientType = "VEGETABLE"
)

func (t *IngredientType) Scan(value interface{}) error {
	*t = IngredientType(value.([]byte))
	return nil
}

func (t IngredientType) Value() (driver.Value, error) {
	return string(t), nil
}

type Ingredient struct {
	gorm.Model
	Name string
	Type IngredientType

	RecipeID uint64
	Recipe   Recipe
}
