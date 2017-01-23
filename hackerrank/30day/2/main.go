/*
Given the meal price (base cost of a meal), tip percent (the percentage of the meal price
being added as tip), and tax percent (the percentage of the meal price being added as tax)
for a meal, find and print the meal's total cost.

Note: Be sure to use precise values for your calculations, or you may end up with an incorrectly rounded result!
*/

package main

import "fmt"

func main() {
	var mealCost float64
	var tipPercent int64
	var taxPercent int64
	var totalCost float64

	fmt.Scanf("%g", &mealCost)
	fmt.Scanf("%d", &tipPercent)
	fmt.Scanf("%d", &taxPercent)

	finalTip := mealCost * (float64(tipPercent) / 100)
	finalTax := mealCost * (float64(taxPercent) / 100)
	totalCost = mealCost + finalTip + finalTax
	fmt.Printf("The total meal cost is %d dollars.\n", RoundToInt32(totalCost))
}

// RoundToInt32 rounds floats into integer numbers.
func RoundToInt32(a float64) int32 {
	if a < 0 {
		return int32(a - 0.5)
	}
	return int32(a + 0.5)
}
