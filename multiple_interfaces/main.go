package main

import "fmt"

// Define the interface
type face1 interface {
	test() int
	test2() string
}

// Struct 1: test
type test struct {
	a int
	b string
}

// Struct 2: another
type another struct {
	x int
	y string
}

// Implementing face1 for *test
func (t *test) test() int {
	return t.a
}

func (t *test) test2() string {
	return t.b
}

// Implementing face1 for *another
func (a *another) test() int {
	return a.x
}

func (a *another) test2() string {
	return a.y
}

// Function that works with the interface
func printInfo(f face1) {
	fmt.Println("test():", f.test())
	fmt.Println("test2():", f.test2())
}

func main() {
	// Create instances
	t1 := &test{a: 10, b: "from test"}
	a1 := &another{x: 99, y: "from another"}

	// Pass to a function accepting interface
	printInfo(t1)
	printInfo(a1)
}
