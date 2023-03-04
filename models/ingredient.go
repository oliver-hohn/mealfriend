package models

import (
	"gorm.io/gorm"
)

type IngredientType string

const (
	BEEF               IngredientType = "beef"
	DAIRY_AND_EGG      IngredientType = "dairy_and_egg"
	FRUIT              IngredientType = "fruit"
	GRAIN_AND_PASTA    IngredientType = "grain_and_pasta"
	LAMB_VEAL_AND_GAME IngredientType = "lamb_veal_and_game"
	LEGUMES            IngredientType = "legumes"
	PORK               IngredientType = "pork"
	POULTRY            IngredientType = "poultry"
	SHELLFISH          IngredientType = "shellfish"
	VEGETABLE          IngredientType = "vegetable"
)

type Ingredient struct {
	gorm.Model
	Name string

	Type IngredientType

	RecipeID uint64
	Recipe   Recipe
}
