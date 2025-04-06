/*
Hands-on exercise #58 - basic funcs
● Hands on exercise

	○ create a func with the identifier foo that returns an int
	○ create a func with the identifier bar that returns an int and a string
	○ call both funcs

○ print out their results
*/
package main

import "fmt"

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "Hello, World!"
}

func main() {

	x := foo()
	fmt.Println(x)

	i, s := bar()
	fmt.Println(i, s)
}
