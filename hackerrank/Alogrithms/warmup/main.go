package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	return (a + b)
}

func main() {
	var a, b, res uint32
	// Reading integer from Stdin
	fmt.Scanf("%v\n%v", &a, &b)
	res = solveMeFirst(a, b)
	fmt.Println(res)
}
