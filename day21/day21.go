package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Food struct {
	Ingredients []string
	Allergens   []string
}

func contains(list []string, item string) bool {
	for _, check := range list {
		if item == check {
			return true
		}
	}
	return false
}

func remove(list []string, find string) []string {
	newList := []string{}
	for _, item := range list {
		if item != find {
			newList = append(newList, item)
		}
	}
	return newList
}

func Parse(input []string) ([]Food, map[string][]string, map[string][]string) {
	foods := []Food{}
	ingredientToAllergen := make(map[string][]string)
	allergenToIngredient := make(map[string][]string)
	for _, line := range input {
		parts := strings.Split(line, " (contains ")
		ingredients := strings.Split(parts[0], " ")
		allergens := strings.Split(parts[1], ", ")
		allergens[len(allergens)-1] = strings.TrimSuffix(allergens[len(allergens)-1], ")")
		foods = append(foods, Food{Ingredients: ingredients, Allergens: allergens})

		for _, ingredient := range ingredients {
			if _, ok := ingredientToAllergen[ingredient]; ok {
				for _, allergen := range allergens {
					if !contains(ingredientToAllergen[ingredient], allergen) {
						ingredientToAllergen[ingredient] = append(ingredientToAllergen[ingredient], allergen)
					}
				}
			} else {
				ingredientToAllergen[ingredient] = allergens
			}

		}

		for _, allergen := range allergens {
			if _, ok := allergenToIngredient[allergen]; ok {
				for _, ingredient := range ingredients {
					if !contains(allergenToIngredient[allergen], ingredient) {
						allergenToIngredient[allergen] = append(allergenToIngredient[allergen], ingredient)
					}
				}
			} else {
				allergenToIngredient[allergen] = ingredients
			}
		}
	}
	return foods, ingredientToAllergen, allergenToIngredient
}

func CountNonAllergens(foods []Food, ingredientToAllergen map[string][]string) (int, []string) {
	nonAllergens := []string{}
	for ingredient, allergens := range ingredientToAllergen {
		isPossibleAllergen := false

		for _, allergen := range allergens {
			isPossibleAllergen = isPossibleAllergen || inEveryFood(foods, allergen, ingredient)
		}

		if !isPossibleAllergen {
			nonAllergens = append(nonAllergens, ingredient)
		}
	}
	count := 0
	for _, food := range foods {
		for _, nonallergen := range nonAllergens {
			if contains(food.Ingredients, nonallergen) {
				count++
			}
		}
	}
	return count, nonAllergens
}

func inEveryFood(foods []Food, allergen string, ingredient string) bool {
	for _, food := range foods {
		if contains(food.Allergens, allergen) {
			if !contains(food.Ingredients, ingredient) {
				return false
			}
		}
	}
	return true
}

func inBoth(list1 []string, list2 []string) []string {
	combined := []string{}
	for _, l1 := range list1 {
		for _, l2 := range list2 {
			if l1 == l2 {
				combined = append(combined, l1)
			}
		}
	}
	return combined
}

func done(rows map[string][]string) bool {
	for _, row := range rows {
		if len(row) > 1 {
			return false
		}
	}
	return true
}

func ComputeAllergens(allergenToIngredient map[string][]string, foods []Food, nonAllergens []string) string {
	possibleAllergenNames := make(map[string][]string)
	for allergen := range allergenToIngredient {
		var possibleIngredientsList [][]string
		for _, food := range foods {
			if contains(food.Allergens, allergen) {
				var possibleIngredients []string
				for _, ingredient := range food.Ingredients {
					if !contains(nonAllergens, ingredient) {
						possibleIngredients = append(possibleIngredients, ingredient)
					}
				}
				possibleIngredientsList = append(possibleIngredientsList, possibleIngredients)
			}
		}
		reduce := possibleIngredientsList[0]
		for _, list := range possibleIngredientsList[1:] {
			reduce = inBoth(reduce, list)
		}
		possibleAllergenNames[allergen] = reduce

	}

	for !done(possibleAllergenNames) {
		for _, row := range possibleAllergenNames {
			if len(row) == 1 {
				// fmt.Println(row[0])
				toRemove := row[0]
				for j, alterRows := range possibleAllergenNames {
					if len(alterRows) > 1 {
						newRow := []string{}
						for _, id := range alterRows {
							if id != toRemove {
								newRow = append(newRow, id)
							}
						}
						possibleAllergenNames[j] = newRow
					}
				}
			}
		}
	}
	allergenList := []string{}
	for allergen := range possibleAllergenNames {
		allergenList = append(allergenList, allergen)
	}
	sort.Strings(allergenList)
	nameList := []string{}
	for _, allergen := range allergenList {
		nameList = append(nameList, possibleAllergenNames[allergen][0])
	}
	return strings.Join(nameList, ",")
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	foods, ingredientToAllergen, allergenToIngredient := Parse(input)
	count, nonAllergens := CountNonAllergens(foods, ingredientToAllergen)
	fmt.Println("Part 1:", count)
	fmt.Println("Part 2::", ComputeAllergens(allergenToIngredient, foods, nonAllergens))
}
