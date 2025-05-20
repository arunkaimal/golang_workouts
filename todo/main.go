package main

import (
	"bufio"
	"fmt"
	"os"
)

// type TodoTask struct {
// 	name string
// }

func main() {
	fmt.Println("Welcome to ToDo List")

	fmt.Println("----------------------------")
	for {
		fmt.Println(`Enter your todo task,If you want to exit please enter "exit"`)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		choice := scanner.Text()

		if choice == "exit" {
			break
		}

		name := scanner.Text()

		fmt.Println(name)
	}
}
