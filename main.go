package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/olekukonko/tablewriter"
)

type IngredientType int

const (
	Meat    = iota
	Chicken = iota
	Fish    = iota
	Other   = iota
)

func (i IngredientType) getName() string {
	switch i {
	case Meat:
		return "Meat"
	case Chicken:
		return "Chicken"
	case Fish:
		return "Fish"
	case Other:
		return "Other"
	default:
		return "N/A"
	}
}

type Ingredient struct {
	Amount uint32
	Name   string
	Type   IngredientType
}

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

type Rule struct {
	Type   IngredientType
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
		fmt.Fprintf(os.Stderr, "Failed to get recipes: %w", err)
		os.Exit(1)
	}

	diet, err := getDiet()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get diet: %w", err)
		os.Exit(1)
	}

	schedule, err := generateSchedule(recipes, diet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get schedule: %w", err)
		os.Exit(1)
	}

	prettyPrintSchedule(schedule)
}

func getRecipes() ([]Recipe, error) {
	recipes := []Recipe{
		{
			Name: "Chicken Pasta Bake",
			Ingredients: []Ingredient{
				{
					Name: "Rigatonni Pasta",
					Type: Other,
				},
				{
					Name: "Chicken Thigs",
					Type: Chicken,
				},
				{
					Name: "Bacon",
					Type: Meat,
				},
			},
		},
		{
			Name: "Lasagna",
			Ingredients: []Ingredient{
				{
					Name: "Beef Mince",
					Type: Meat,
				},
				{
					Name: "Lasagne Sheets",
					Type: Other,
				},
			},
		},
		{
			Name: "Quesadillas",
			Ingredients: []Ingredient{
				{
					Name: "Beef Mince",
					Type: Meat,
				},
				{
					Name: "Tortillas",
					Type: Other,
				},
			},
		},
		{
			Name: "Moqueca",
			Ingredients: []Ingredient{
				{
					Name: "White Fish (Halibut, Black Cod, Sea Bass, Cod)",
					Type: Fish,
				},
				{
					Name: "Coconut Milk",
					Type: Other,
				},
			},
		},
		{
			Name: "Fish Risotto",
			Ingredients: []Ingredient{
				{
					Name: "Prawns",
					Type: Fish,
				},
				{
					Name: "Mussels",
					Type: Fish,
				},
				{
					Name: "Risotto Rice",
					Type: Other,
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
				Type:   Meat,
				Amount: 1,
			},
			{
				Type:   Fish,
				Amount: 1,
			},
			{
				Type:   Chicken,
				Amount: 1,
			},
		},
	}

	return diet, nil
}

func generateSchedule(recipes []Recipe, diet *Diet) (map[string]*Recipe, error) {
	recipesByDay := map[string]*Recipe{}

	for i, rule := range diet.Rules {
		recipe, err := getRandomRecipeWith(rule.Type, recipes)
		if err != nil {
			return nil, err
		}

		recipesByDay[fmt.Sprintf("day_%v", i+1)] = recipe
	}

	return recipesByDay, nil
}

func getRandomRecipeWith(i IngredientType, r []Recipe) (*Recipe, error) {
	rand.Shuffle(len(r), func(a, b int) { r[a], r[b] = r[b], r[a] })

	for _, recipe := range r {
		if recipe.HasIngredient(i) {
			return &recipe, nil
		}
	}

	return nil, fmt.Errorf("unable to find Recipe with Ingredient: %v", i.getName())
}

type byDay [][]string

func (d byDay) Len() int      { return len(d) }
func (d byDay) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d byDay) Less(i, j int) bool {
	return d[i][0] < d[j][0]
}

func prettyPrintSchedule(s map[string]*Recipe) {
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
