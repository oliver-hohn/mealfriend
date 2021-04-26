package models

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

type Ingredient struct {
	Amount uint64
	Name   string
	Type   IngredientType
}
