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
	arun.GetStatus()
	deferred()
}

func (u User) GetStatus() {
	fmt.Println("Is status active:", u.Status)
	fmt.Println("His Email is:", u.Email)
	fmt.Println("His Name is:", u.Name)
}

func deferred() {
	a := 10
	b := 5

	defer fmt.Println("hai iam deffered and i will only execute in last")
	fmt.Println(a + b)
}
