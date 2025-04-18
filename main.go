/*
Hands-on exercise #72 - closure

Closure is when we have “enclosed” the scope of a variable in some code block. For this
hands-on exercise, create a func which “encloses” the scope of a variable
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	g := incrementor()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

	h := powinator(2)
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func powinator(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}
