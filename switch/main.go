package main

import "fmt"

func main() {
	day := 6

	switch day {

	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknowm")

	}

	grade := 70

	switch {
	case grade >= 90 && grade <= 100:
		fmt.Println("Grade A")
	case grade >= 80 && grade < 90:
		fmt.Println("Grade B")
	default:
		fmt.Println("Fail")
	}
}
