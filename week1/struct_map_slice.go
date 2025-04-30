package main

import (
	"fmt"
	"strings"
)

// define a struct
type Contact struct {
	Name  string
	Email string
	Age   int
}

func Pointer() {
	var a int = 5
	var p *int = &a // p points to a

	fmt.Println("a =", a)   // 5
	fmt.Println("p =", p)   // address (e.g., 0xc000014098)
	fmt.Println("*p =", *p) // 5 (dereferencing p)

	// Change value through pointer
	*p = 10
	fmt.Println("After changing *p:")
	fmt.Println("a =", a) // 10
}

func SliceDemo() {
	nums := []int{1, 2, 3}
	nums = append(nums, 4) // add element
	fmt.Println("Slice contents:", nums)
	fmt.Println("Second element:", nums[1])
}

func MapDemo() {
	capitals := make(map[string]string)
	capitals["USA"] = "Washington, D.C."
	capitals["France"] = "Paris"
	fmt.Println("Map contents:", capitals)
	fmt.Println("Capital of France:", capitals["France"])

	// Check if key exists
	val, ok := capitals["Germany"]
	if ok {
		fmt.Println("Capital of Germany:", val)
	} else {
		fmt.Println("Germany not found in map.")
	}
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)     // split string into words
	counts := make(map[string]int) // create a map to store counts

	for _, word := range words {
		counts[word]++ // increment count for each word
	}
	fmt.Println("Word counts:", counts)
	return counts
}

func StructDemo() {
	c1 := Contact{Name: "Alice", Email: "alice@example.com", Age: 30}
	c2 := Contact{"Bob", "bob@example.com", 25}
	fmt.Println("Contact 1:", c1)
	fmt.Println("Contact 2 Name:", c2.Name)

	// Slice of structs
	contacts := []Contact{c1, c2}
	fmt.Println("\nAll contacts:")
	for _, c := range contacts {
		fmt.Printf("- %s (%d years old)\n", c.Name, c.Age)
	}
}
