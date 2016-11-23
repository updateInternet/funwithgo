package main

import (
	"fmt"
	"time"
)

// mergeSort divides the array into chunks
func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	mid := len(a) / 2

	var left = mergeSort(a[:mid])
	var right = mergeSort(a[mid:])
	fmt.Println(merge(left, right))
	return merge(left, right)
}

// merge left and right slice into a newly created slice
func merge(left, right []int) []int {

	output := make([]int, len(left)+len(right))

	var i = 0
	var j = 0

	for k := 0; k < len(output); k++ {
		if i >= len(left) {
			output[k] = right[j]
			j++
			continue
		} else if j >= len(right) {
			output[k] = left[i]
			i++
			continue
		}
		if left[i] > right[j] {
			output[k] = right[j]
			j++
		} else {
			output[k] = left[i]
			i++
		}
	}
	// for i < len(left) && j < len(right) {
	// 	if left[i] <= right[j] {
	// 		output[i+j] = left[i]
	// 		i++
	// 	} else {
	// 		output[i+j] = right[j]
	// 		j++
	// 	}
	// }
	//
	// for i < len(left) {
	// 	output[i+j] = left[i]
	// 	i++
	// }
	// for j < len(right) {
	// 	output[i+j] = right[j]
	// 	j++
	// }

	return output
}

func main() {
	random := []int{76, 3, 45, 12, 11, 98, 23, 67, 9037, 456, 678}
	fmt.Println("Unsorted Array:", random)
	startTime := time.Now()
	random = mergeSort(random)
	endTime := time.Now()
	fmt.Println("Sorted Array:", random)
	// Print elapsed time in nanoSeconds
	fmt.Println("Time: ", endTime.UnixNano()-startTime.UnixNano())
}
