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
