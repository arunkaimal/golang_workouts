package main

import "fmt"

func main() {
	// var name *string

	// fmt.Println(name)

	num := 10

	newNum := &num

	
	
	fmt.Println("Name is: ", newNum)
	fmt.Println("Name is :", *newNum)
	
	*newNum = *newNum *2

	fmt.Println(num)

}
