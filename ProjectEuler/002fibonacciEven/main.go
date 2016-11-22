/* Problem 2  - Even Fibonacci numbers
https://projecteuler.net/problem=2

Each new term in the Fibonacci sequence is generated by adding the previous two terms.
By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four
million, find the sum of the even-valued terms.

The formula: enter image description here
*/

package main

import "fmt"

func fibonacci(limit int) (total int) {
	total = 0
	// implement Fibonacci formula inside loop
	for a, b := 0, 1; a < limit; a, b = b, a+b {
		// other way if math.Mod(float64(a), 2)
		if a%2 == 0 {
			// Sum only even numbers
			total += a
		}
	}
	return
}

func main() {
	fmt.Println(fibonacci(100)) //  first 100 number
	fmt.Println(fibonacci(4e6)) //  4 milllion numbers
}