package main

import "fmt"

type Person struct {
	Name string
}

func (person Person) getHello() string {
	return "Hello " + person.Name
}
func main() {
	dhoniaridho := Person{
		Name: "dhoniaridho",
	}
	fmt.Println(dhoniaridho.getHello())
}
