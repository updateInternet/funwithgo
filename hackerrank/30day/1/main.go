/*
Declare  variables: one of type int, one of type double, and one of type String.
Read  lines of input from stdin (according to the sequence given in the Input Format section below) and initialize your  variables.
Use the  operator to perform the following operations:
Print the sum of  plus your int variable on a new line.
Print the sum of  plus your double variable to a scale of one decimal place on a new line.
Concatenate  with the string you read as input and print the result on a new line.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i uint32 = 4
	var d float32 = 4.0
	var s = "HackerRank "

	scanner := bufio.NewReader(os.Stdin)
	var a uint32
	var b float32

	fmt.Scanf("%d", &a)
	fmt.Scanf("%g", &b)
	input, _ := scanner.ReadString('\n')
	fmt.Println(a + i)
	fmt.Printf("%.1f \n", float32(b+d))
	fmt.Println(s + input)
}

/*
  var i uint32 = 4
	var d float32 = 4.0

	var a uint32
	var b float32
	var c bytes.Buffer
	fmt.Scanf("%d", &a)
	fmt.Scanf("%g", &b)
	c.Write([]byte("HackerRank "))
	fmt.Println(a + i)
	fmt.Printf("%.1f \n", float32(b+d))
	fmt.Fprintf(&c, "is the best place to learn and practice coding!")
	c.WriteTo(os.Stdout)
*/

/*
func main() {
	var i uint32 = 4
	var d float32 = 4.0

	var a uint32
	var b float32
	var c = "HackerRank "
	fmt.Scanf("%d", &a)
	fmt.Scanf("%g", &b)
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	fmt.Println(a + i)
	fmt.Printf("%.1f \n", float32(b+d))
	e := scan.Text()
	fmt.Println(c + e)
}

*/
