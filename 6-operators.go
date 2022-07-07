package main

import "fmt"

func main() {

	a := 20 + 20 // Adds two operands
	b := 20 - 20 // Subtracts second operand from the first
	c := 20 * 20 // Multiplies both operands
	d := 20 / 20 // Divides the numerator by the denominator.
	e := 20 % 20 // Modulus operator; gives the remainder after an integer division.
	f := 20
	f++ // Increment operator. It increases the integer value by one.
	g := 20
	g-- // Decrement operator. It decreases the integer value by one.

	fmt.Println("Operators")
	fmt.Println("Add 20 + 20", a)
	fmt.Println("Subtracts 20-20", b)
	fmt.Println("Multiplies 20*20", c)
	fmt.Println("Divides 20/20", d)
	fmt.Println("Divides 20%20", e)
	fmt.Println("Increment from 20 +1", f)
	fmt.Println("Decrement from 20 -1", g)

	fmt.Println("==", a == a)                    // 	It checks if the values of two operands are equal or not; if yes, the condition becomes true.
	fmt.Println("!=", a != a)                    // 	It checks if the values of two operands are equal or not; if the values are not equal, then the condition becomes true.
	fmt.Println("greater than or equal", a >= f) // It checks if the value of the left operand is greater than or equal to the value of the right operand; if yes, the condition becomes true.
	fmt.Println("less than or equal", a <= f)    // 	It checks if the value of left operand is less than or equal to the value of right operand; if yes, the condition becomes true.
	fmt.Println("greater than", a > a)
	fmt.Println("less than", a < a)

	fmt.Println("AND", true && false) // 	Called Logical AND operator. If both the operands are false, then the condition becomes false.
	fmt.Println("OR", true || false)  // 		Called Logical OR Operator. If any of the two operands is true, then the condition becomes true.
	fmt.Println("NOT", !true)         // 	Called Logical NOT Operator. Use to reverses the logical state of its operand. If a condition is true, then Logical NOT operator will make it false.

}
