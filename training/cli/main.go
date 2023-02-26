package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/oliver-hohn/mealfriend/models"
)

var foodFilePath = flag.String("food_file_path", "", "path to seed food file")
var foodCategoryFilePath = flag.String("food_category_file_path", "", "path to food category file path")
var outputPath = flag.String("output_path", "training/data/labelled.jsonl", "path to output labelled data")

type food struct {
	Description string
	CategoryId  int
}

type category struct {
	Id          int
	Description string
}

func main() {
	flag.Parse()

	if *foodFilePath == "" {
		log.Fatal("missing food_file_path")
	}

	if *foodCategoryFilePath == "" {
		log.Fatal("missing food_category_file_path")
	}

	if *outputPath == "" {
		log.Fatal("missing output_path")
	}

	categoriesById, err := buildCategories(*foodCategoryFilePath)
	if err != nil {
		log.Fatal(err)
	}

	foods, err := buildFoods(*foodFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if err := export(foods, categoriesById, *outputPath); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Exported labelled data to: %s\n", *outputPath)
}

func buildCategories(path string) (map[int]*category, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read category file: %w", err)
	}

	ret := map[int]*category{}

	categoryCsv := csv.NewReader(file)
	rowCount := 0
	for {
		rowCount += 1
		record, err := categoryCsv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("unable to read category row: %w", err)
		}

		if rowCount == 1 {
			// Ignore the first row
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("unable to parse ID: %s (on row: %d), because: %w", record[0], rowCount, err)
		}

		ret[id] = &category{Id: id, Description: record[2]}
	}

	return ret, nil
}

func buildFoods(path string) ([]*food, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read food file: %w", err)
	}

	ret := []*food{}

	foodCsv := csv.NewReader(file)
	rowCount := 0
	for {
		rowCount += 1
		record, err := foodCsv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("unable to read food row: %w", err)
		}

		if rowCount == 1 {
			// Ignore the first row
			continue
		}

		category := record[3]
		if len(category) == 0 {
			ret = append(ret, &food{Description: record[2]})
			continue
		}

		categoryId, err := strconv.Atoi(category)
		if err != nil {
			return nil, fmt.Errorf("unable to parse category ID: %s (on row: %d), because: %w", record[3], rowCount, err)
		}

		ret = append(ret, &food{Description: record[2], CategoryId: categoryId})
	}

	return ret, nil
}

type gcpClassification struct {
	ClassificationAnnotation struct {
		DisplayName string `json:"displayName"`
	} `json:"classificationAnnotation"`
	TextContent            string `json:"textContent"`
	DataItemResourceLabels struct {
		AiPlatformTag string `json:"aiplatform.googleapis.com/ml_use""`
	} `json:"dataItemResourceLabels"`
}

func export(foods []*food, categoriesById map[int]*category, outputPath string) error {
	classifications := []*gcpClassification{}

	for _, food := range foods {
		c := gcpClassification{TextContent: food.Description}

		var fmtCategory string

		if rawCategory, ok := categoriesById[food.CategoryId]; ok {
			fmtCategory = generateFormattedCategory(rawCategory.Description)
		} else {
			fmtCategory = "none_of_the_above"
		}

		c.ClassificationAnnotation.DisplayName = fmtCategory
		c.DataItemResourceLabels.AiPlatformTag = "training"

		classifications = append(classifications, &c)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("unable to create output file: %w", err)
	}

	w := bufio.NewWriter(f)

	for _, c := range classifications {
		b, err := json.Marshal(c)
		if err != nil {
			return fmt.Errorf("unable to serialize: %v, because: %w", c, err)
		}

		if _, err := w.Write(b); err != nil {
			return fmt.Errorf("unable to write: %v, to file because: %w", c, err)
		}

		if _, err := w.WriteString("\n"); err != nil {
			return fmt.Errorf("unable to write new line: %w", err)
		}
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("unable to flush writer: %w", err)
	}

	return nil
}

func generateFormattedCategory(rawCategory string) string {
	switch rawCategory {
	case "Dairy and Egg Products":
		return models.DAIRY_AND_EGG
	case "Poultry Products":
		return models.POULTRY
	case "Fruits and Fruit Juices":
		return models.FRUIT
	case "Pork Products":
		return models.PORK
	case "Vegetables and Vegetable Products":
		return models.VEGETABLE
	case "Beef Products":
		return models.BEEF
	case "Finfish and Shellfish Products":
		return models.SHELLFISH
	case "Legumes and Legume Products":
		return models.LEGUMES
	case "Lamb, Veal, and Game Products":
		return models.LAMB_VEAL_AND_GAME
	case "Cereal Grains and Pasta":
		return models.GRAIN_AND_PASTA
	default:
		return "none_of_the_above"
	}
}
