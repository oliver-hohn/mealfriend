package classification

import (
	"regexp"

	"github.com/oliver-hohn/mealfriend/models"
)

var classificationRules = []Rule{
	{Condition: regexp.MustCompile("beef"), IngredientType: models.BEEF},

	{Condition: regexp.MustCompile("butter"), IngredientType: models.DAIRY_AND_EGG},
	{Condition: regexp.MustCompile("cheese"), IngredientType: models.DAIRY_AND_EGG},
	{Condition: regexp.MustCompile(`(double|heavy|single)\s*cream`), IngredientType: models.DAIRY_AND_EGG},
	{Condition: regexp.MustCompile("egg"), IngredientType: models.DAIRY_AND_EGG},
	{Condition: regexp.MustCompile("milk"), IngredientType: models.DAIRY_AND_EGG},

	{Condition: regexp.MustCompile("apple"), IngredientType: models.FRUIT},
	{Condition: regexp.MustCompile("banana"), IngredientType: models.FRUIT},
	{Condition: regexp.MustCompile("coconut"), IngredientType: models.FRUIT},
	{Condition: regexp.MustCompile("lemon"), IngredientType: models.FRUIT},
	{Condition: regexp.MustCompile("orange"), IngredientType: models.FRUIT},
	{Condition: regexp.MustCompile("tomato"), IngredientType: models.FRUIT},

	{Condition: regexp.MustCompile("gnocchi"), IngredientType: models.GRAIN_AND_PASTA},
	{Condition: regexp.MustCompile("orzo"), IngredientType: models.GRAIN_AND_PASTA},
	{Condition: regexp.MustCompile("pasta"), IngredientType: models.GRAIN_AND_PASTA},
	{Condition: regexp.MustCompile("rice"), IngredientType: models.GRAIN_AND_PASTA},
	{Condition: regexp.MustCompile("spaghetti"), IngredientType: models.GRAIN_AND_PASTA},

	{Condition: regexp.MustCompile("lamb"), IngredientType: models.LAMB_VEAL_AND_GAME},

	{Condition: regexp.MustCompile("chickpea"), IngredientType: models.LEGUMES},
	{Condition: regexp.MustCompile("lentils"), IngredientType: models.LEGUMES},
	{Condition: regexp.MustCompile(`(red\s*kidney|black|white)\s*bean`), IngredientType: models.LEGUMES},

	{Condition: regexp.MustCompile("pork"), IngredientType: models.PORK},

	{Condition: regexp.MustCompile("chicken"), IngredientType: models.POULTRY},
	{Condition: regexp.MustCompile("turkey"), IngredientType: models.POULTRY},

	{Condition: regexp.MustCompile("cod"), IngredientType: models.FISH},
	{Condition: regexp.MustCompile("halibut"), IngredientType: models.FISH},
	{Condition: regexp.MustCompile("salmon"), IngredientType: models.FISH},
	{Condition: regexp.MustCompile(`sea\s*bass`), IngredientType: models.FISH},
	{Condition: regexp.MustCompile("tuna"), IngredientType: models.FISH},

	{Condition: regexp.MustCompile("clams"), IngredientType: models.SHELLFISH},
	{Condition: regexp.MustCompile("crab"), IngredientType: models.SHELLFISH},
	{Condition: regexp.MustCompile("lobster"), IngredientType: models.SHELLFISH},
	{Condition: regexp.MustCompile("mussels"), IngredientType: models.SHELLFISH},
	{Condition: regexp.MustCompile("oysters"), IngredientType: models.SHELLFISH},
	{Condition: regexp.MustCompile("scallops"), IngredientType: models.SHELLFISH},
	{Condition: regexp.MustCompile("shrimp"), IngredientType: models.SHELLFISH},

	{Condition: regexp.MustCompile("aubergine"), IngredientType: models.VEGETABLE},
	{Condition: regexp.MustCompile(`(bell|red|green|orange)\s*pepper`), IngredientType: models.VEGETABLE},
	{Condition: regexp.MustCompile("broccoli"), IngredientType: models.VEGETABLE},
	{Condition: regexp.MustCompile("carrot"), IngredientType: models.VEGETABLE},
	{Condition: regexp.MustCompile("celery"), IngredientType: models.VEGETABLE},
	{Condition: regexp.MustCompile("corn"), IngredientType: models.VEGETABLE},
	{Condition: regexp.MustCompile("(courgette|zucchini)"), IngredientType: models.VEGETABLE},
	{Condition: regexp.MustCompile(`green\s*bean`), IngredientType: models.VEGETABLE},
	{Condition: regexp.MustCompile("mushroom"), IngredientType: models.VEGETABLE},
	{Condition: regexp.MustCompile("pea"), IngredientType: models.VEGETABLE},
	{Condition: regexp.MustCompile("spinach"), IngredientType: models.VEGETABLE},
}

func Classify(raw string) models.IngredientType {
	for _, r := range classificationRules {
		if r.Condition.MatchString(raw) {
			return r.IngredientType
		}
	}

	return models.UNKNOWN
}
