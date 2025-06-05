package main

import (
	"fmt"
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

	//MINIMUM STEPS TO REACH TARGET POSITION USING BFS
	stepsWithBFS := knight.knightmoves(target,8)
	fmt.Printf("Minimum steps to reach target (%d, %d): %d\n", target.X, target.Y, stepsWithBFS)

}

func (k *Knight) knightmoves(target Position, size int) int {
	possibleMovementOnX := []int{1, -1, 2, -2, 2, -2, 1, -1}
	possibleMovementOnY := []int{2, -2, 1, -1, -1, 1, -2, 2}

	queue := []Position{k.Position}
	mapVisited := make(map[Position]bool)
	mapVisited[k.Position] = true
	steps := 0

	if k.Position==target {
		return steps
	}

	for len(queue) > 0 {
		steps++
		queueLen:=len(queue)
		for range queueLen {
			currentPosition := queue[0]
			queue = queue[1:]
			for j := range 8 {
				newX := currentPosition.X + possibleMovementOnX[j]
				newY := currentPosition.Y + possibleMovementOnY[j]
				newPosition := Position{X: newX, Y: newY}
				if newX == target.X && newY == target.Y {
					return steps
				} else if isValidPosition(newPosition, size, &mapVisited) {
					mapVisited[newPosition] = true
					queue = append(queue, newPosition)
				}

			}
		}
	}
	return steps
}

func isValidPosition(position Position, size int, mapVisited *map[Position]bool) bool {
	// Check if position is within bounds
	if position.X < size && position.Y < size && position.X >= 0 && position.Y >= 0 {
		// Check if the position has not been visited
		if !(*mapVisited)[position] {
			return true
		}
	}
	return false
}
