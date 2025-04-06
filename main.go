/*
Hands-on exercise #63 - tests in go #1
*/

package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(Add(3, 4))
}
