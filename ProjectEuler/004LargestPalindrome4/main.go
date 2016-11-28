/*Problem 4
A palindromic number reads the same both ways. The largest palindrome made from the
product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import "fmt"

func main() {
	largestPalindrome := 0
	a := 100
	for ; a <= 999; a = a + 1 {
		b := 100
		for ; b <= 999; b = b + 1 {
			if c := a * b; isPalindrome(c) && c > largestPalindrome {
				largestPalindrome = c
			}
		}
	}
	fmt.Println(largestPalindrome)
}

func reversed(n int) int {
	reversed := 0
	for ; n > 0; n = n / 10 {
		reversed = 10*reversed + n%10
	}
	return reversed
}

// check if its the number passing into func is same when reversed
func isPalindrome(n int) bool {
	return n == reversed(n)
}
