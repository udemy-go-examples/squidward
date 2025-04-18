/*
Hands-on exercise #70 - func return

	● Create a func
		○ which returns a func
	■ which returns 42
	● assign the returned func to a variable
	● call the returned func
	● print
*/
package main

import "fmt"

func main() {

	xf := myfunc()
	fmt.Println(xf())
}

// returns a function
func myfunc() func() int {
	return func() int {
		return 42
	}
}
