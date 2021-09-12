package utils

import (
	"strings"

	pbmodels "github.com/oliver-hohn/mealfriend/protos/models"
)

func NewIngredient(raw string) *pbmodels.Ingredient {
	return &pbmodels.Ingredient{Name: strings.TrimSpace(raw)}
}
