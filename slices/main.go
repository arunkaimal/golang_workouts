package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 10, 8, 4}

	var index  = 3

	nums = append(nums[:index],nums[index+1:]...)

	fmt.Println(nums)

}
