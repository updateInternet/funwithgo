/*To complete this challenge, you must save a line of input from stdin to a variable,
print Hello, World. on a single line, and finally print the value of your variable on a second line.
*/

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
