package main

import "fmt"

func main() {

	type person struct {
		name string
		age  int
	}

	var people = person{}

	people.name = "Ahmad Ridhoni"
	people.age = 20

	fmt.Println(people)
}
