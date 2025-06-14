package main

import "fmt"

type Vehicle interface {
	Start()
}

type Car struct {
	Brand string
}

type Bike struct {
	Brand string
}

func (c *Car) Start() {
	fmt.Println("Car", c.Brand, "is starting with a key 🔑")
}

func (b *Bike) Start() {
	fmt.Println("Bike", b.Brand, "is kick-starting 🦵")
}

// Car-specific method
func (c *Car) PlayMusic() {
	fmt.Println("Car is playing music 🎵")
}

