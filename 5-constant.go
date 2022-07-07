package main

import "fmt"

func main() {

	const WIDTH int = 200  // Naming constant better capitalize
	const HEIGHT int = 200 // Naming constant better capitalize

	var area int

	area = WIDTH * HEIGHT

	fmt.Println("Constant", WIDTH, HEIGHT, area)
}
