package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	input := []string{
		"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
		"trh fvjkl sbzzf mxmxvkd (contains dairy)",
		"sqjhc fvjkl (contains soy)",
		"sqjhc mxmxvkd sbzzf (contains fish)",
	}

	foods, ingredientToAllergen, allergenToIngredient := Parse(input)

	fmt.Println(foods)
	fmt.Println(ingredientToAllergen)
	fmt.Println(allergenToIngredient)
}

func TestFindNonAllergens(t *testing.T) {
	input := []string{
		"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
		"trh fvjkl sbzzf mxmxvkd (contains dairy)",
		"sqjhc fvjkl (contains soy)",
		"sqjhc mxmxvkd sbzzf (contains fish)",
	}

	foods, ingredientToAllergen, _ := Parse(input)
	count, _ := CountNonAllergens(foods, ingredientToAllergen)

	require.Equal(t, 5, count)
}

func TestWorkOutAllergen(t *testing.T) {
	input := []string{
		"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
		"trh fvjkl sbzzf mxmxvkd (contains dairy)",
		"sqjhc fvjkl (contains soy)",
		"sqjhc mxmxvkd sbzzf (contains fish)",
	}

	foods, ingredientToAllergen, allergenToIngredient := Parse(input)
	_, nonAllergens := CountNonAllergens(foods, ingredientToAllergen)

	require.Equal(t, "mxmxvkd,sqjhc,fvjkl", ComputeAllergens(allergenToIngredient, foods, nonAllergens))
}
