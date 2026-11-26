package main

import (
	"fmt"

	"example.com/app/mathutil"
)

func main() {
	var a int = 7
	var b int = 11

	// Factorial of a=7
	factorial := mathutil.Factorial(a)
	fmt.Printf("Factorial of %d is: %d\n", a, factorial)

	// Power of a=7 raised to b=11
	power := mathutil.Power(a, b)
	fmt.Printf("%d raised to the power of %d is: %d\n", a, b, power)

	// Reverse a string
	str := "Muskan"
	reversedStr := mathutil.Stringreverse(str)
	fmt.Printf("Reversed string of '%s' is: '%s'\n", str, reversedStr)

	// Count vowels in a string
	vowelCount := mathutil.CountVowels(str)
	fmt.Printf("Number of vowels in '%s' is: %d\n", str, vowelCount)

}
