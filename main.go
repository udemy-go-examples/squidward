/*
Hands-on exercise #64 - tests in go #2 - unit tests

create a unit test for these three functions
*/

package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)
	x := doMath(42, 16, add)
	fmt.Println(x)
	y := doMath(42, 16, subtract)
	fmt.Println(y)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}
