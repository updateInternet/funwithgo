// Implement Merge Sort in go - phase 1
// Use go routines to sort - phase 2
// Use distributed systems logic to implement merge-sort - phase 3

package main

import "fmt"

func main() {
	unsorted := []int{11, 57, 14, 6, 11, 0, 1, 12, 15, 100, 16, 2, 4, 16, 32, 8}
	fmt.Println(unsorted)
	l, r := split(unsorted)
	fmt.Println(l, r)
	fmt.Println(merge(l, r))
}

// split array into two
func split(unsorted []int) ([]int, []int) {
	return unsorted[0 : len(unsorted)/2], unsorted[len(unsorted)/2:] // Split into 2 halves
}

func merge(a []int, b []int) []int {
	arr := make([]int, len(a)+len(b))
	fmt.Println(arr)
	return arr
}
