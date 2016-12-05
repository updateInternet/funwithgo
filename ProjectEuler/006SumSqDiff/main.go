/*Problem 6 - Sum square difference

The sum of the squares of the first ten natural numbers is  1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is  (1 + 2 + ... + 10)^2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and
the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural
numbers and the square of the sum.
*/

package main

import "fmt"

// N is the defined constant
const N = 100

func main() {
	ans := N * (N + 1) / 2 // 10 * (10+1)/2 = 55; 10 * 11/2; 110/2
	ans *= ans
	sos := 0
	for i := 0; i <= N; i++ {
		sos += i * i
	}
	ans -= sos
	fmt.Println(ans)
}
