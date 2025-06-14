package main

import (
	"fmt"
)

// Interface
type Animal interface {
	Speak()
}

// Structs
type Dog struct{}
type Cat struct{}

// Interface implementations
//Dog implements Animal interface
//Cat implements Animal interface
// Both Dog and Cat have a Speak method
// that satisfies the Animal interface.
// This allows us to use polymorphism
// to treat different types of animals uniformly.
// They can be used interchangeably
// as long as they implement the Speak method.
// This is an example of polymorphism in Go.
func (d *Dog) Speak() {
	fmt.Println("buff buff")
}

func (c *Cat) Speak() {
	fmt.Println("meow meow")
}

// Dog-specific method
func (d *Dog) Cuddle() {
	fmt.Println("Dog is cuddling 🐶❤️")
}

// Helper function to call Speak
func describeAnimal(a Animal) {
	a.Speak()
}

func main() {
	animals := []Animal{&Dog{}, &Cat{}}

	for _, a := range animals {
		describeAnimal(a)

		// Type assertion to check if a is *Dog
		if dog, ok := a.(*Dog); ok {
			dog.Cuddle()
		}
	}
}
