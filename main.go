/*
Hands-on exercise #60 - defer func
	● “defer” multiple functions in main
		○ show that a deferred func runs after the func containing it exits.
		○ determine the order in which the multiple defer funcs run
*/

package main

import "fmt"

func main() {

	// deferred function run LIFO order
	fmt.Println("Start of main")
	defer fmt.Println("This is the first deferred function call")
	defer fmt.Println("This is the second deferred function call")
	defer fmt.Println("This is the third deferred function call")
	defer fmt.Println("This is the fourth deferred function call")

	fmt.Println("End of main")
}
