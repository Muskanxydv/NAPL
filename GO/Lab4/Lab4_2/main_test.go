package main

import (
	"testing"
)

func TestDivide(t *testing.T) {
	result, err := divide(10, 2)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != 5 {
		t.Errorf("Expected 5, got %f", result)
	}

	_, err = divide(10, 0)

	if err == nil {
		t.Error("Expected division by zero error")
	}
}

func TestSquareRoot(t *testing.T) {

	result, err := squareRoot(25)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != 5 {
		t.Errorf("Expected 5, got %f", result)
	}

	_, err = squareRoot(-9)

	if err == nil {
		t.Error("Expected negative number error")
	}
}

func TestCheckInput(t *testing.T) {
	result, err := checkInput("Hello")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != "Hello" {
		t.Errorf("Expected Hello, got %s", result)
	}

	_, err = checkInput("")

	if err == nil {
		t.Error("Expected empty input error")
	}
}
