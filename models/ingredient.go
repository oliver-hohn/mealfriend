package models

type IngredientType int

const (
	Meat    = iota
	Chicken = iota
	Fish    = iota
	Other   = iota
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
	Amount uint32
	Name   string
	Type   IngredientType
}
