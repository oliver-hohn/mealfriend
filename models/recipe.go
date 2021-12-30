package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Code string
	Name string

	Ingredients []Ingredient
}
