/*Problem 5 - Smallest multiple
2520 is the smallest number that can be divided by each of the numbers from 1 to 10
without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from
1 to 20?
*/
package main

import "fmt"

func main() {
	var a, i int64 = 1, 2
	for ; i <= 20; i++ {
		a = lcm(a, i)
	}
	fmt.Println(a)
}

// least common multiple
func lcm(a, b int64) int64 {
	x := a * b / gcd(a, b)
	return x
}

//greatest common divisor
func gcd(a, b int64) int64 {
	c := a % b
	if c == 0 {
		return b
	}
	return gcd(b, c)
}
