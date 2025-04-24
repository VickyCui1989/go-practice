// main.go
package main

import (
	"fmt"
	// "math"
)

func main() {
	fmt.Println("🚀 Day 1 - Go is working!")

	name := "XXX"
	age := 20
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	fmt.Println("Loop from 1 to 5:")
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	greet(name)
}

func greet(name string) {
	fmt.Printf("Hello, %s! Welcome to Go!\n", name)
}
