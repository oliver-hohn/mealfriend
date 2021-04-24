package models

type Recipe struct {
	Name        string
	Ingredients []Ingredient
}

func (r *Recipe) HasIngredient(i IngredientType) bool {
	for _, ingredient := range r.Ingredients {
		if ingredient.Type == i {
			return true
		}
	}

	return false
}
