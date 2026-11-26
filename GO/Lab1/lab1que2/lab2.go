package main

import "fmt"

func main() {
	var choice int

	for {
		fmt.Println("Calculator Menu")
		fmt.Println("1) Integer operation")
		fmt.Println("2) Float operation")
		fmt.Println("3) Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var a, b int

			fmt.Print("\nEnter first integer: ")
			fmt.Scan(&a)

			fmt.Print("Enter second integer: ")
			fmt.Scan(&b)

			fmt.Println("Addition:", a+b)
			fmt.Println("Subtraction:", a-b)
			fmt.Println("Multiplication:", a*b)

		case 2:
			var a, b float32

			fmt.Print("Enter first float: ")
			fmt.Scan(&a)

			fmt.Print("Enter second float: ")
			fmt.Scan(&b)

			fmt.Println("Addition:", a+b)
			fmt.Println("Subtraction:", a-b)
			fmt.Println("Multiplication:", a*b)

		case 3:
			fmt.Println("Exiting calculator. Bbyee!")
			return

		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}
