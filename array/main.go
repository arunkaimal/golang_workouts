package main

import "fmt"

func main() {
	// var nums = [4]string{"Arun", "Anu", "Manu", "Rohith"}

	// for i := 0; i < len(nums); i++ {

	// 	fmt.Println(i)
	// }

	// var names [10]string

	// fmt.Printf("%q \n", names)

	// var numbers = [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	// fmt.Println(len(numbers))

	originalArray := [5]int{1, 2, 3, 4, 5}
	coppiedArray := originalArray

	coppiedArray[0] = 100
	fmt.Println(originalArray)
	fmt.Println(coppiedArray)
}
