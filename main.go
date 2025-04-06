/*
Hands-on exercise #61 - method
	● Create a user defined struct with
		○ the identifier “person”
		○ the fields:
			■ first
			■ age
	● attach a method to type person with
		○ the identifier “speak”
		○ the method should have the person say their name and age
	● create a value of type person
	● call the method from the value of type person
*/

package main

import "fmt"

type Person struct {
	first string
	age   int
}

func (p Person) speak() {
	fmt.Printf("My name is %v and my age is %v years old\n", p.first, p.age)
}

func main() {

	p := Person{
		first: "James",
		age:   42,
	}

	p.speak()
}
