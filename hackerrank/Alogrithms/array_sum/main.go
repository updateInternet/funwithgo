/*Given an array of  integers, can you find the sum of its elements? */

package main

import "fmt"

func main() {
	var a, b int
	sum := 0
	// Reading integer from Stdin
	fmt.Scanf("%v", &a)
	for i := 1; i <= a; i++ {
		fmt.Scanf("%v", &b)
		sum += b
	}
	fmt.Println(sum)
}
