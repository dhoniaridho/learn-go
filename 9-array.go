package main

import "fmt"

func main() {
	var arr = [6]int{
		56,
		45,
		45,
		4545,
		54,
		45,
	}
	var arrString = [2]string{"hello", "world"}
	var names [2]string

	names[0] = "Doni"
	names[1] = "dhoniaridho"

	for _, item := range names {
		fmt.Println(item)
	}

	for _, item := range arrString {
		fmt.Println(item)
	}

	for item, index := range arr {
		fmt.Println(item, index)
	}
}
