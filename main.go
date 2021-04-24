package main

import (
	"fmt"
	"main/models"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Rule struct {
	Type   models.IngredientType
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

func getRecipes() ([]models.Recipe, error) {
	recipes := []models.Recipe{
		{
			Name: "Chicken Pasta Bake",
			Ingredients: []models.Ingredient{
				{
					Name: "Rigatonni Pasta",
					Type: models.Other,
				},
				{
					Name: "Chicken Thigs",
					Type: models.Chicken,
				},
				{
					Name: "Bacon",
					Type: models.Meat,
				},
			},
		},
		{
			Name: "Lasagna",
			Ingredients: []models.Ingredient{
				{
					Name: "Beef Mince",
					Type: models.Meat,
				},
				{
					Name: "Lasagne Sheets",
					Type: models.Other,
				},
			},
		},
		{
			Name: "Quesadillas",
			Ingredients: []models.Ingredient{
				{
					Name: "Beef Mince",
					Type: models.Meat,
				},
				{
					Name: "Tortillas",
					Type: models.Other,
				},
			},
		},
		{
			Name: "Moqueca",
			Ingredients: []models.Ingredient{
				{
					Name: "White Fish (Halibut, Black Cod, Sea Bass, Cod)",
					Type: models.Fish,
				},
				{
					Name: "Coconut Milk",
					Type: models.Other,
				},
			},
		},
		{
			Name: "Fish Risotto",
			Ingredients: []models.Ingredient{
				{
					Name: "Prawns",
					Type: models.Fish,
				},
				{
					Name: "Mussels",
					Type: models.Fish,
				},
				{
					Name: "Risotto Rice",
					Type: models.Other,
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
				Type:   models.Meat,
				Amount: 1,
			},
			{
				Type:   models.Fish,
				Amount: 1,
			},
			{
				Type:   models.Chicken,
				Amount: 1,
			},
		},
	}

	return diet, nil
}

func generateSchedule(recipes []models.Recipe, diet *Diet) (map[string]*models.Recipe, error) {
	recipesByDay := map[string]*models.Recipe{}

	for i, rule := range diet.Rules {
		recipe, err := getRandomRecipeWith(rule.Type, recipes)
		if err != nil {
			return nil, err
		}

		recipesByDay[fmt.Sprintf("day_%v", i+1)] = recipe
	}

	return recipesByDay, nil
}

func getRandomRecipeWith(i models.IngredientType, r []models.Recipe) (*models.Recipe, error) {
	rand.Shuffle(len(r), func(a, b int) { r[a], r[b] = r[b], r[a] })

	for _, recipe := range r {
		if recipe.HasIngredient(i) {
			return &recipe, nil
		}
	}

	return nil, fmt.Errorf("unable to find Recipe with Ingredient: %v", i.GetName())
}

type byDay [][]string

func (d byDay) Len() int      { return len(d) }
func (d byDay) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d byDay) Less(i, j int) bool {
	return d[i][0] < d[j][0]
}

func prettyPrintSchedule(s map[string]*models.Recipe) {
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
