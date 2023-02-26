package models

import (
	"gorm.io/gorm"
)

const (
	BEEF               = "beef"
	DAIRY_AND_EGG      = "dairy_and_egg"
	FRUIT              = "fruit"
	GRAIN_AND_PASTA    = "grain_and_pasta"
	LAMB_VEAL_AND_GAME = "lamb_veal_and_game"
	LEGUMES            = "legumes"
	PORK               = "pork"
	POULTRY            = "poultry"
	SHELLFISH          = "shellfish"
	VEGETABLE          = "vegetable"
)

type Ingredient struct {
	gorm.Model
	Name string

	RecipeID uint64
	Recipe   Recipe
}
