package main

import "fmt"

type Man struct {
	Name string
}

/**
* without pointer
* by default pass by value in go
* and value wont changed while not using pointer
**/
func (man Man) say() {
	man.Name = "Mr. " + man.Name
}

// with pointer
func (man *Man) people() {
	man.Name = "Mr. " + man.Name
}

func main() {
	dhoniaridho := Man{Name: "dhoniaridho"}
	ridho := Man{Name: "ridho"}
	dhoniaridho.people()
	ridho.say()
	fmt.Println(dhoniaridho.Name)
	fmt.Println(ridho.Name)
}
