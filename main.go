/*
Hands-on exercise #69 - func expression

	● Assign a func to a variable, then call that func
*/
package main

import "fmt"

func main() {

	xf := func(value int) string {
		return fmt.Sprintf("Anonymous function run with value %v", value)
	}

	fmt.Println(xf(69))
}
