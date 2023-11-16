package helpers

import (
	"fmt"

	"github.com/oliver-hohn/mealfriend/models"
)

func PrettyPrintRecipe(r *models.Recipe) {
	fmt.Printf("Title: %s\n", r.Name)
	fmt.Printf("Source: %s\n", r.Source.String())
	fmt.Printf("Image: %s\n", r.ImageURL())
	fmt.Printf("Tags: %v\n", r.Tags)
	fmt.Printf("Cook time: %s\n", r.CookTime.String())

	fmt.Println("Ingredients:")
	for _, i := range r.Ingredients {
		fmt.Printf("\t- %s\n", i)
	}
}
