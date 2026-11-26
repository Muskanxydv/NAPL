package main

import "fmt"

func main() {
	fmt.Printf("Hello, Muskan!")
	//two float point variables
	var a float32 = 7.1
	var b float32 = 1.9
	//two integer variables
	var c int = 7
	var d int = 17

	fmt.Printf("Float Operations:\n")
	//addition of the two float point numbers defined before
	var sum float32 = a + b
	fmt.Printf("Sum of two float point numbers:%g + %g = %g\n", a, b, sum)
	//subtraction of the two float point numbers defined before
	var sub float32 = a - b
	fmt.Printf("Subtraction of two float point numbers:%g - %g = %g\n", a, b, sub)
	//multiplication of the two float point numbers defined before
	var mul float32 = a * b
	fmt.Printf("Multiplication of two float point numbers:%g * %g = %g\n", a, b, mul)

	fmt.Printf("Integer Operations:\n")
	//addition of the two integer numbers defined before
	var sum1 int = c + d
	fmt.Printf("Sum of two integer numbers:%d + %d = %d\n", c, d, sum1)
	//subtraction of the two integer numbers defined before
	var sub1 int = c - d
	fmt.Printf("Subtraction of two integer numbers:%d - %d = %d\n", c, d, sub1)
	//multiplication of the two integer numbers defined before
	var mul1 int = c * d
	fmt.Printf("Multiplication of two integer numbers:%d * %d = %d\n", c, d, mul1)
}
