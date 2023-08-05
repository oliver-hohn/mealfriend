package tagger

import (
	"regexp"
	"strings"

	"github.com/oliver-hohn/mealfriend/models"
)

type ingredientTagRule struct {
	Condition *regexp.Regexp
	Tag       models.Tag
}

var ingredientTagRules = []ingredientTagRule{
	{Condition: regexp.MustCompile("beef"), Tag: models.BEEF},

	{Condition: regexp.MustCompile("butter"), Tag: models.DAIRY},
	{Condition: regexp.MustCompile("cheddar"), Tag: models.DAIRY},
	{Condition: regexp.MustCompile("cheese"), Tag: models.DAIRY},
	{Condition: regexp.MustCompile(`(double|heavy|single|sour)\s*cream`), Tag: models.DAIRY},
	{Condition: regexp.MustCompile("milk"), Tag: models.DAIRY},
	{Condition: regexp.MustCompile("mozzarella"), Tag: models.DAIRY},

	{Condition: regexp.MustCompile("egg"), Tag: models.EGG},

	{Condition: regexp.MustCompile("cod"), Tag: models.FISH},
	{Condition: regexp.MustCompile("halibut"), Tag: models.FISH},
	{Condition: regexp.MustCompile("salmon"), Tag: models.FISH},
	{Condition: regexp.MustCompile(`sea\s*bass`), Tag: models.FISH},
	{Condition: regexp.MustCompile("tuna"), Tag: models.FISH},

	{Condition: regexp.MustCompile("apple"), Tag: models.FRUIT},
	{Condition: regexp.MustCompile("avocado"), Tag: models.FRUIT},
	{Condition: regexp.MustCompile("banana"), Tag: models.FRUIT},
	{Condition: regexp.MustCompile("coconut"), Tag: models.FRUIT},
	{Condition: regexp.MustCompile("lemon"), Tag: models.FRUIT},
	{Condition: regexp.MustCompile("orange"), Tag: models.FRUIT},
	{Condition: regexp.MustCompile("tomato"), Tag: models.FRUIT},

	{Condition: regexp.MustCompile("rice"), Tag: models.GRAIN},

	{Condition: regexp.MustCompile("chickpea"), Tag: models.LEGUMES},
	{Condition: regexp.MustCompile("lentils"), Tag: models.LEGUMES},
	{Condition: regexp.MustCompile(`(red\s*kidney|black|white)\s*bean`), Tag: models.LEGUMES},

	{Condition: regexp.MustCompile("gnocchi"), Tag: models.PASTA},
	{Condition: regexp.MustCompile("orzo"), Tag: models.PASTA},
	{Condition: regexp.MustCompile("pasta"), Tag: models.PASTA},
	{Condition: regexp.MustCompile("spaghetti"), Tag: models.PASTA},

	{Condition: regexp.MustCompile("pork"), Tag: models.PORK},

	{Condition: regexp.MustCompile("chicken"), Tag: models.POULTRY},
	{Condition: regexp.MustCompile("turkey"), Tag: models.POULTRY},

	{Condition: regexp.MustCompile("clams"), Tag: models.SHELLFISH},
	{Condition: regexp.MustCompile("crab"), Tag: models.SHELLFISH},
	{Condition: regexp.MustCompile("lobster"), Tag: models.SHELLFISH},
	{Condition: regexp.MustCompile("mussels"), Tag: models.SHELLFISH},
	{Condition: regexp.MustCompile("oysters"), Tag: models.SHELLFISH},
	{Condition: regexp.MustCompile("scallops"), Tag: models.SHELLFISH},
	{Condition: regexp.MustCompile("shrimp"), Tag: models.SHELLFISH},

	{Condition: regexp.MustCompile("aubergine"), Tag: models.VEGETABLE},
	{Condition: regexp.MustCompile(`(bell|red|green|orange)\s*pepper`), Tag: models.VEGETABLE},
	{Condition: regexp.MustCompile("broccoli"), Tag: models.VEGETABLE},
	{Condition: regexp.MustCompile("carrot"), Tag: models.VEGETABLE},
	{Condition: regexp.MustCompile("celery"), Tag: models.VEGETABLE},
	{Condition: regexp.MustCompile("corn"), Tag: models.VEGETABLE},
	{Condition: regexp.MustCompile("(courgette|zucchini)"), Tag: models.VEGETABLE},
	{Condition: regexp.MustCompile(`green\s*bean`), Tag: models.VEGETABLE},
	{Condition: regexp.MustCompile("mushroom"), Tag: models.VEGETABLE},
	{Condition: regexp.MustCompile("pea"), Tag: models.VEGETABLE},
	{Condition: regexp.MustCompile("potato"), Tag: models.VEGETABLE},
	{Condition: regexp.MustCompile("spinach"), Tag: models.VEGETABLE},
}

func tagForIngredient(raw string) models.Tag {
	r := strings.ToLower(raw)

	for _, rule := range ingredientTagRules {
		if rule.Condition.MatchString(r) {
			return rule.Tag
		}
	}

	return models.UNSPECIFIED
}
