package utils

import (
	"strings"

	"github.com/oliver-hohn/mealfriend/classification"
	"github.com/oliver-hohn/mealfriend/models"
)

func NewIngredient(raw string) models.Ingredient {
	return models.Ingredient{Name: strings.TrimSpace(raw), Type: classification.Classify(raw)}
}
