package tagger

import "github.com/oliver-hohn/mealfriend/models"

func TagsForRecipe(r *models.Recipe) []models.Tag {
	tagHistory := map[models.Tag]bool{}

	ret := []models.Tag{}
	for _, i := range r.Ingredients {
		t := tagForIngredient(i)

		if t == models.UNSPECIFIED {
			continue
		}

		if _, ok := tagHistory[t]; !ok {
			tagHistory[t] = true
			ret = append(ret, t)
		}
	}

	return ret
}
