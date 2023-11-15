package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/oliver-hohn/mealfriend/models"
	"github.com/oliver-hohn/mealfriend/tagger"
)

func NewRecipe(name string, ingredients []string, source *url.URL) *models.Recipe {
	r := &models.Recipe{
		Name:        name,
		Code:        strcase.ToCamel(name),
		Ingredients: ingredients,
		Source:      source,
	}

	r.Tags = tagger.TagsForRecipe(r)

	return r
}

var hoursRegex = regexp.MustCompile(`(\d+)\s*(hours?|hrs?)`)
var minutesRegex = regexp.MustCompile(`(\d+)\s*(minutes?|mins?)`)

type NoDurationError struct {
	Message string
}

func (e *NoDurationError) Error() string {
	return e.Message
}

func NewDuration(duration string) (time.Duration, error) {
	var err error

	var hours int
	// First match is always the section of the string that matches the regex (e.g. "1 hour"),
	// the second match is the decimal captured by the regex (e.g. "1")
	if matches := hoursRegex.FindStringSubmatch(duration); len(matches) > 1 {
		hours, err = strconv.Atoi(matches[1])
		if err != nil {
			return -1, fmt.Errorf("unable to convert %s into an int: %w", matches[1], err)
		}
	}

	var minutes int
	if matches := minutesRegex.FindStringSubmatch(duration); len(matches) > 1 {
		minutes, err = strconv.Atoi(matches[1])
		if err != nil {
			return -1, fmt.Errorf("unable to convert %s into an int: %w", matches[1], err)
		}
	}

	if hours == 0 && minutes == 0 {
		return -1, &NoDurationError{Message: fmt.Sprintf("unable to parse %s into a duration", duration)}
	}

	return time.ParseDuration(fmt.Sprintf("%dh%dm", hours, minutes))
}
