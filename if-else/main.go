package main

import "fmt"

func main() {
	grade := 55

	if grade >= 90 && grade <= 100 {
		fmt.Println("Grade A")
	} else if grade >= 80 && grade < 90 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Fail")
	}
}
