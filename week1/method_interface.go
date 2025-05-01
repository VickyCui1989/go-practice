package main

import "fmt"

// Method with value receiver
func (c Contact) Greet() {
	fmt.Printf("Hi, I'm %s. You can reach me at %s.\n", c.Name, c.Email)
}

// Method with pointer receiver (can modify original)
func (c *Contact) HaveBirthday() {
	c.Age++
}

// Define an interface
type Speaker interface {
	Greet()
}

func MethodInterfaceDemo() {
	// Create a Contact
	c := Contact{Name: "Charlie", Email: "charlie@example.com", Age: 28}

	// Call methods
	c.Greet()        // Method call
	c.HaveBirthday() // Modify age using pointer receiver
	fmt.Println("After birthday, age is:", c.Age)

	// Interface usage
	var s Speaker = c
	s.Greet() // Can call because Contact implements Greet()

	// Slice of Speaker interface
	people := []Speaker{
		Contact{Name: "Dana", Email: "dana@example.com", Age: 35},
		Contact{Name: "Eli", Email: "eli@example.com", Age: 40},
	}

	fmt.Println("\nSpeakers:")
	for _, p := range people {
		p.Greet() // Interface method call
	}
}
