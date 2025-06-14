package main

import (
	"fmt"
)

// Define the interface
type face1 interface {
	test() int
	test2() string
}

// Struct 'test' implements the interface
type _test struct {
	a int
	b string
}

// Methods with pointer receivers
func (t *_test) test() int {
	return t.a
}

func (t *_test) test2() string {
	return t.b
}

// A function that accepts any type implementing face1
func printInfo(f face1) {
	fmt.Println("Value of a:", f.test())
	fmt.Println("Value of b:", f.test2())
}

type Animal interface {
	Speak()
}

type Dog struct {
	Animal
}

func (d *Dog) Speak() {
	fmt.Println("buff buff")
}





func main() {
	// Create an instance of test and pass to the interface-based function
	x := &_test{a: 1, b: "test value"}

	// This function doesn't care about the actual struct,
	// just that it satisfies the interface
	fmt.Println(x.test())
	printInfo(x)

	d := &Dog{}
	d.Speak()
	fmt.Println("Hello, World!")
	

}
