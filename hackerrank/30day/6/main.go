/*
Given a string, , of length  that is indexed from  to , print its even-indexed and odd-indexed
characters as  space-separated strings on a single line (see the Sample below for more detail).

Note:  is considered to be an even index.
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var num int
	fmt.Scan(&num)

	scanner := bufio.NewReader(os.Stdin)
	var lines []string
	if num > 0 && num <= 10 {
		for i := 0; i < num; i++ {
			line, _ := scanner.ReadString('\n')
			lines = append(lines, line)
		}
	} else {
		fmt.Println("Enter a number between 1 and 10")
	}

	for _, v := range lines {
		if len(v) > 1 {
			var (
				fword, lword string
			)
			for i := range v {
				if v[i] != '\r' && v[i] != '\n' {
					if math.Mod(float64(i), 2) == 0 {
						fword = fword + string(v[i])
					} else {
						lword = lword + string(v[i])
					}
				}
			}
			fmt.Printf("%s %s\n", fword, lword)
		}
	}
}
