package main

import (
	"fmt"
	"math"
)

type Position struct {
	X int
	Y int
}

type Knight struct {
	Position Position
}

func main() {
	knight := Knight{Position: Position{X: 1, Y: 3}}
	target := Position{X: 5, Y: 0}

	// minimum steps to reach the target position
	steps := knight.minimumStepsToReach(target)
	fmt.Printf("Minimum steps to reach target (%d, %d): %d\n", target.X, target.Y, steps)
}

func (k *Knight) minimumStepsToReach(target Position) int {
	//min steps required to reach the target position
	dx := math.Abs(float64(k.Position.X - target.X))
	dy := math.Abs(float64(k.Position.Y - target.Y))
	if dx == 0 && dy == 0 {
		return 0
	}
	if dx == 1 && dy == 2 || dx == 2 && dy == 1 {
		return 1
	}
	if dx == 2 && dy == 2 {
		return 4
	}
	if dx == 1 && dy == 1 {
		return 3
	}

	return 0

}
