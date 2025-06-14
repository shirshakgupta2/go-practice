package main

import "fmt"

// Interfaces
type Walker interface {
	Walk()
}

type Barker interface {
	Bark()
}

type Charger interface {
	Charge()
}

// Struct that implements all 3 interfaces
type RobotDog struct {
	Name string
}

func (r *RobotDog) Walk() {
	fmt.Println(r.Name, "is walking.")
}

func (r *RobotDog) Bark() {
	fmt.Println(r.Name, "says: Buff Buff (robotically).")
}

func (r *RobotDog) Charge() {
	fmt.Println(r.Name, "is charging 🔋.")
}

// Interface Composition
type SmartPet interface {
	Walker
	Barker
	Charger
}

func performActions(p SmartPet) {
	p.Walk()
	p.Bark()
	p.Charge()
}

