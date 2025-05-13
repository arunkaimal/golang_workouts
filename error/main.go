package main

import (
	"errors"
	"fmt"
)

func division(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("couldn't Divisible")
	}
	return num1 / num2, nil
}

func main() {
	answer, err := division(20, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Result :", answer)
}
