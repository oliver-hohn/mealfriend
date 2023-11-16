package models

import (
	"net/url"
	"time"
)

type Tag string

const (
	BEEF      Tag = "beef"
	DAIRY     Tag = "dairy"
	EGG       Tag = "egg"
	FISH      Tag = "fish"
	FRUIT     Tag = "fruit"
	GRAIN     Tag = "grain"
	LEGUMES   Tag = "legumes"
	PASTA     Tag = "pasta"
	PORK      Tag = "pork"
	POULTRY   Tag = "poultry"
	SHELLFISH Tag = "shellfish"
	VEGETABLE Tag = "vegetable"

	UNSPECIFIED Tag = "unspecified"
)

var AvailableTags = []Tag{
	BEEF,
	DAIRY,
	EGG,
	FISH,
	FRUIT,
	GRAIN,
	LEGUMES,
	PASTA,
	PORK,
	POULTRY,
	SHELLFISH,
	VEGETABLE,
}

type Recipe struct {
	Code string
	Name string

	Source *url.URL

	Ingredients []string
	Tags        []Tag

	CookTime time.Duration
	Image    *url.URL
}

func (r *Recipe) CookTimeStr() string {
	if r.CookTime == 0 {
		return ""
	}
	return r.CookTime.String()
}

func (r *Recipe) ImageURL() string {
	if r.Image == nil {
		return ""
	}
	return r.Image.String()
}
