package main

import "fmt"

func main() {
	a, b := readarrays()
	fmt.Println(compareint(a, b))
}

func compareint(a, b []int) (int, int) {
	var alice, bob int
	for i := 0; i < 3; i++ {
		if a[i]|b[i] >= 100 {
			break
		}
		if a[i] > b[i] {
			alice++
		} else if b[i] > a[i] {
			bob++
		}
	}
	return alice, bob
}

func readarrays() ([]int, []int) {
	a := make([]int, 3)
	for i := range a {
		fmt.Scanf("%d", &a[i])
	}
	b := make([]int, 3)
	for i := range b {
		fmt.Scanf("%d", &b[i])
	}
	return a, b
}
