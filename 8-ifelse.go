package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		evenOrOdd(i)
	}
}

func evenOrOdd(input int) int {
	if input%2 == 0 {
		fmt.Println(input, "Even")
	} else {
		fmt.Println(input, "Odd")
	}
	return input
}
