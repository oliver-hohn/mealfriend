package classification

import (
	"regexp"

	"github.com/oliver-hohn/mealfriend/models"
)

type Rule struct {
	Condition      *regexp.Regexp
	IngredientType models.IngredientType
}
