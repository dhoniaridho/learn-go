package main

import "fmt"

func main() {
	var x int = 4
	var y *int = &x
	var z int = x

	change(&x, 1)
	fmt.Println("x", x)
	fmt.Println("x", &x)
	fmt.Println("y", *y)
	fmt.Println("y", y)
	fmt.Println("z", z)
	fmt.Println("z", &z)
}

func change(original *int, value int) {
	*original = value
}
