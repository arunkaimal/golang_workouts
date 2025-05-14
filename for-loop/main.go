package main

import "fmt"

func main() {

	number := []int{11, 2, 3, 4, 5}

	for ind, value := range number {
		fmt.Printf("index : %d & Value : %d\n", ind, value)
	}

	for _, value := range number {
		fmt.Printf("Value : %d\n", value)
	}
}
