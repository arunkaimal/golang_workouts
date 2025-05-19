package main

import "fmt"

func main() {
	namePlaces := make(map[string]string)

	namePlaces["Arun"] = "Thrissur"
	namePlaces["Anu"] = "Kottayam"
	namePlaces["Vyshnav"] = "Thrissur"

	for key, val := range namePlaces {
		fmt.Printf("Key is %s and Value is %s\n", key, val)
	}

	cities := map[string]string {
		"Kerala":"Thrissur",
		"Maharashtra":"Mumbai",
	}

	for key, val := range cities {
		fmt.Printf("---------Key is %s and Value is %s\n", key, val)
	}
	
}
