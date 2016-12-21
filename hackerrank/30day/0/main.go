package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World.")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	fmt.Println(scan.Text())
}
