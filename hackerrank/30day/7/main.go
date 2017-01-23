package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var num int
	fmt.Scan(&num)
	if num > 0 && num <= 1000 {
		scanner := bufio.NewReader(os.Stdin)
		buffer, _ := scanner.ReadString('\n')
		buffer = strings.Trim(buffer, "\r\n")
		bufferSlice := strings.Split(buffer, " ")
		// reverse
		for i := num - 1; i >= 0; i-- {
			fmt.Printf("%s ", bufferSlice[i])
		}
	} else {
		fmt.Println("Enter a number between 1 and 1000")
	}
}
