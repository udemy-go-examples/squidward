/*
Hands-on exercise #71 - callback
A “callback” is when we pass a func into a func as an argument. For this exercise,

	● pass a func into a func as an argument
		○ func square(n int) int
		○ printSquare(f func(int)int, int) string
*/
package main

import "fmt"

type callbackFunc func(int) int

func square(n int) int {
	return n * n
}

func printSquare(f callbackFunc, value int) string {
	return fmt.Sprintf("Square of %v is %v", value, f(value))
}

func main() {

	fmt.Println(printSquare(square, 9))
}
