package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Whats ur name?")
	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n')
	fmt.Println("My name is", name)
}
