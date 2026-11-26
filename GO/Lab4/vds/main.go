package main

import (
	"fmt"
	"math"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("negative number not allowed")
	}
	return math.Sqrt(n), nil
}

func checkInput(s string) (string, error) {
	if s == "" {
		return "", fmt.Errorf("input cannot be empty")
	}
	return s, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Division:", result)
	}

	root, err := squareRoot(-9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Square Root:", root)
	}

	input, err := checkInput("")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Input:", input)
	}
}
