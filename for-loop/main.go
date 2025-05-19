package main

import "fmt"

func main() {

	// number := []int{11, 2, 3, 4, 5}

	// for ind, value := range number {
	// 	fmt.Printf("index : %d & Value : %d\n", ind, value)
	// }

	// for _, value := range number {
	// 	fmt.Printf("Value : %d\n", value)
	// }

	var nums = []int{15, 25, 35, 14, 20, 85, 20}

	odd := []int{}
	even := []int{}
	for i := range nums {
		if i%2 == 0 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}
	}

	fmt.Println("Even Numbers Are :", even)
	fmt.Println("Odd Numbers Are :", odd)
}
