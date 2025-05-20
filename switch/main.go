package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var operation string

	fmt.Print("Enter 1st number: ")
	fmt.Scanln(&num1)

	fmt.Println("Enter Operation (+,-,*,/):")
	fmt.Scanln(&operation)

	fmt.Print("Enter 2nd number: ")
	fmt.Scanln(&num2)

	switch operation {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	}
}
