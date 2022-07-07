package main

import "fmt"

func main() {
	/**
	* unsigned number positive UINT
	* signed number can be negative INT
	 */

	// Signed Variables can use negative value
	var a int = -10       // int auto by operating system example windows 64 bit default be 64bit
	var b int8 = -10      // Signed 8-bit integers (-128 to 127)
	var c int16 = -4000   // Signed 16-bit integers (-32768 to 32767)
	var d int32 = -40000  // Signed 32-bit integers (-2147483648 to 2147483647)
	var e int64 = -500000 // Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// Unsigned Variables must positive value
	var f uint = 10       // uint auto by operating system example windows 64 bit default be 64bit
	var g uint8 = 10      // Unsigned 8-bit integers (0 to 255)
	var h uint16 = 4000   // Unsigned 16-bit integers (0 to 65535)
	var i uint32 = 40000  // Unsigned 32-bit integers (0 to 4294967295)
	var j uint64 = 500000 // Unsigned 64-bit integers (0 to 18446744073709551615)

	// Boolean Value True or False
	var k bool = false
	var l bool = true

	var m float32 = 343.5                   // IEEE-754 32-bit floating-point numbers
	var n float64 = -65656565656565656343.5 // IEEE-754 64-bit floating-point numbers

	var o complex64 = 3343434             // Complex numbers with float32 real and imaginary parts
	var p complex128 = 945804958408054850 // Complex numbers with float64 real and imaginary parts

	fmt.Println("Signed Variables", a, b, c, d, e)
	fmt.Println("Unsigned Variables", f, g, h, i, j)
	fmt.Println("Boolean", k, l)
	fmt.Println("Float", m, n, o, p)
}
