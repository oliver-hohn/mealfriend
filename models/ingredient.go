package models

import (
	"main/utils"
	"regexp"
	"strings"
)

type IngredientType int

const (
	Meat = iota + 1
	Chicken
	Fish
	Other
)

func (i IngredientType) GetName() string {
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

type Quantity struct {
	Amount string
	Unit   string
}

type Ingredient struct {
	Quantity Quantity
	Name     string
	Type     IngredientType
}

var quantityRegex = regexp.MustCompile(
	`(?i)(?P<amount>\d+\s*/\s*\d+|\d+\s*\-\s*\d+|\d+)\s*(?P<unit>g|mg|kg|oz|cup(s)?|leg(s)?)?(?P<description>.*)`, // hell-regex
)

var quantityIndexByGroup = utils.MustMapGroupToIndex(quantityRegex)

func NewIngredient(raw string) Ingredient {
	matches := quantityRegex.FindStringSubmatch(raw)
	if len(matches) == 0 {
		return Ingredient{Name: strings.TrimSpace(raw)}
	}

	quantity := Quantity{
		Amount: strings.TrimSpace(matches[quantityIndexByGroup["amount"]]),
		Unit:   strings.TrimSpace(matches[quantityIndexByGroup["unit"]]),
	}

	return Ingredient{
		Quantity: quantity,
		Name:     strings.TrimSpace(matches[quantityIndexByGroup["description"]]),
	}
}
