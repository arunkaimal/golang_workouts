package main

import (
	"bufio"
	"fmt"
	"os"
)

type TodoTask struct {
	name string
}

var todoList []TodoTask

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

		todo := TodoTask{name}
		todoList = append(todoList, todo)

		fmt.Println(todoList)
	}

	for _, todo := range todoList {
		fmt.Println(todo.name)
	}
}
