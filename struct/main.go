package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Struct")
	arun := User{"Arun", "arun@gmail.com", true, 25}
	fmt.Println(arun)
	fmt.Printf("Arun's details are: %+v\n", arun)
	fmt.Printf("Name is : %v\n", arun.Name)
}
