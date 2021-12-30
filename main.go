package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/olekukonko/tablewriter"
	pbmodels "github.com/oliver-hohn/mealfriend/protos/models"
)

type Rule struct {
	Type   pbmodels.Ingredient_Type
	Amount uint32
}

type Diet struct {
	Rules        []Rule
	NumberOfDays uint32
}

func main() {
	rand.Seed(time.Now().UnixNano())
	recipes, err := getRecipes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get recipes: %v", err)
		os.Exit(1)
	}

	diet, err := getDiet()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get diet: %v", err)
		os.Exit(1)
	}

	schedule, err := generateSchedule(recipes, diet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get schedule: %v", err)
		os.Exit(1)
	}

	prettyPrintSchedule(schedule)
}

func getRecipes() ([]*pbmodels.Recipe, error) {
	recipes := []*pbmodels.Recipe{
		{
			Name: "Chicken Pasta Bake",
			Ingredients: []*pbmodels.Ingredient{
				{
					Name: "Rigatonni Pasta",
					Type: pbmodels.Ingredient_UNSPECIFIED,
				},
				{
					Name: "Chicken Thigs",
					Type: pbmodels.Ingredient_CHICKEN,
				},
				{
					Name: "Bacon",
					Type: pbmodels.Ingredient_PORK,
				},
			},
		},
		{
			Name: "Lasagna",
			Ingredients: []*pbmodels.Ingredient{
				{
					Name: "Beef Mince",
					Type: pbmodels.Ingredient_BEEF,
				},
				{
					Name: "Lasagne Sheets",
					Type: pbmodels.Ingredient_UNSPECIFIED,
				},
			},
		},
		{
			Name: "Quesadillas",
			Ingredients: []*pbmodels.Ingredient{
				{
					Name: "Beef Mince",
					Type: pbmodels.Ingredient_BEEF,
				},
				{
					Name: "Tortillas",
					Type: pbmodels.Ingredient_UNSPECIFIED,
				},
			},
		},
		{
			Name: "Moqueca",
			Ingredients: []*pbmodels.Ingredient{
				{
					Name: "White Fish (Halibut, Black Cod, Sea Bass, Cod)",
					Type: pbmodels.Ingredient_FISH,
				},
				{
					Name: "Coconut Milk",
					Type: pbmodels.Ingredient_UNSPECIFIED,
				},
			},
		},
		{
			Name: "Fish Risotto",
			Ingredients: []*pbmodels.Ingredient{
				{
					Name: "Prawns",
					Type: pbmodels.Ingredient_FISH,
				},
				{
					Name: "Mussels",
					Type: pbmodels.Ingredient_FISH,
				},
				{
					Name: "Risotto Rice",
					Type: pbmodels.Ingredient_UNSPECIFIED,
				},
			},
		},
	}

	return recipes, nil
}

func getDiet() (*Diet, error) {
	diet := &Diet{
		NumberOfDays: 3,
		Rules: []Rule{
			{
				Type:   pbmodels.Ingredient_BEEF,
				Amount: 1,
			},
			{
				Type:   pbmodels.Ingredient_FISH,
				Amount: 1,
			},
			{
				Type:   pbmodels.Ingredient_CHICKEN,
				Amount: 1,
			},
		},
	}

	return diet, nil
}

func generateSchedule(recipes []*pbmodels.Recipe, diet *Diet) (map[string]*pbmodels.Recipe, error) {
	recipesByDay := map[string]*pbmodels.Recipe{}

	for i, rule := range diet.Rules {
		recipe, err := getRandomRecipeWith(rule.Type, recipes)
		if err != nil {
			return nil, err
		}

		recipesByDay[fmt.Sprintf("day_%v", i+1)] = recipe
	}

	return recipesByDay, nil
}

func recipeHasIngredient(r *pbmodels.Recipe, i pbmodels.Ingredient_Type) bool {
	for _, ingredient := range r.GetIngredients() {
		if ingredient.GetType() == i {
			return true
		}
	}

	return false
}

func getRandomRecipeWith(i pbmodels.Ingredient_Type, r []*pbmodels.Recipe) (*pbmodels.Recipe, error) {
	rand.Shuffle(len(r), func(a, b int) { r[a], r[b] = r[b], r[a] })

	for _, recipe := range r {
		if recipeHasIngredient(recipe, i) {
			return recipe, nil
		}
	}

	return nil, fmt.Errorf("unable to find Recipe with Ingredient: %v", i.Descriptor().FullName())
}

type byDay [][]string

func (d byDay) Len() int      { return len(d) }
func (d byDay) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d byDay) Less(i, j int) bool {
	return d[i][0] < d[j][0]
}

func prettyPrintSchedule(s map[string]*pbmodels.Recipe) {
	data := [][]string{}

	for day, recipe := range s {
		data = append(data, []string{
			day, recipe.Name,
		})
	}

	sort.Sort(byDay(data))

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Day", "Recipe"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
