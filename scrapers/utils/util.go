package utils

import (
	"net/url"

	"github.com/iancoleman/strcase"
	"github.com/oliver-hohn/mealfriend/models"
	"github.com/oliver-hohn/mealfriend/tagger"
)

func NewRecipe(name string, ingredients []string, source *url.URL) *models.Recipe {
	r := &models.Recipe{
		Name:        name,
		Code:        strcase.ToCamel(name),
		Ingredients: ingredients,
		Source:      source,
	}

	r.Tags = tagger.TagsForRecipe(r)

	return r
}
