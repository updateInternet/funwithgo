package main

import "fmt"

func main() {
	var num, result int
	// tables to run against
	fmt.Scanf("%d", &num)
	if num >= 2 && num <= 20 {
		for i := 1; i < 11; i++ {
			result = num * i
			fmt.Printf("%d x %d = %d \n", num, i, result)
		}
	} else {
		fmt.Println("Enter a number between 2 and 20")
	}
}
