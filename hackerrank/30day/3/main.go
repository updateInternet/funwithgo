package main

import "fmt"

func main() {
	var num int
	fmt.Scanf("%d", &num)
	// number should be less than 100
	if num <= 100 {
		// odd
		if num%2 != 0 {
			fmt.Println("Weird")
		} else {
			// even
			if num < 6 && num > 1 {
				fmt.Println("Not Weird")
			}
			if num < 21 && num > 5 {
				fmt.Println("Weird")
			}
			if num > 20 {
				fmt.Println("Not Weird")
			}
		}
	}
}

/*
var num int
fmt.Scanf("%d", &num)
if num <= 100 {
	if num%2 == 0 {
				fmt.Println("Not Weird")
	} else {
		fmt.Println("Weird")
	}
} else {
	fmt.Println("Enter num less than 100")
}
*/

/*
var num int
fmt.Scanf("%d", &num)
if num <= 100 {
	if num%2 != 0 {
		fmt.Println("Weird")
	} else if num := 2; num < 5 {
		fmt.Println("Not Weird")
	} else if num := 6; num < 20 {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
} else {
	fmt.Println("Enter number less than 100")
}
*/
