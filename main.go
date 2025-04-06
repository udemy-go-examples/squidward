/*
Hands-on exercise #59 - variadic func
	● create a func with the identifier foo that
		○ takes in a variadic parameter of type int
		○ pass in a value of type []int into your func (unfurl the []int)
		○ returns the sum of all values of type int passed in
	● create a func with the identifier bar that
		○ takes in a parameter of type []int
		○ returns the sum of all values of type int passed in
*/

package main

import "fmt"

func foo(ii ...int) int {
	total := 0
	for _, i := range ii {
		total += i
	}

	return total
}

func bar(ii []int) int {
	total := 0
	for _, i := range ii {
		total += i
	}

	return total
}

func main() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	x := foo(xi...)
	fmt.Println(x)

	y := bar(xi)
	fmt.Println(y)
}
