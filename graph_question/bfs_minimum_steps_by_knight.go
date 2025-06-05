package graphquestion


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


	//MINIMUM STEPS TO REACH TARGET POSITION USING BFS
	stepsWithBFS := knight.minimumStepsToReachBFS(target)
	fmt.Printf("Minimum steps to reach target (%d, %d): %d\n", target.X, target.Y, stepsWithBFS)


}

func (k *Knight) minimumStepsToReachBFS(target Position) int {
	// Knight's possible moves
	// Each move is represented as a change in x and y coordinates
	// The knight can move in an "L" shape: two squares in one direction and one square perpendicular
	// to that direction.
	// The possible moves are:
	// (-2, 1), (-1, 2), (1, 2), (2, 1), (2, -1), (1, -2), (-1, -2), (-2, -1)
	dx := []int{-2, -1, 1, 2, 2, 1, -1, -2}
	dy := []int{1, 2, 2, 1, -1, -2, -2, -1}

	queue := []Position{k.Position}
	visited := make(map[Position]bool)
	visited[k.Position] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			current := queue[i]
			if current == target {
				return steps
			}
			// Explore all possible knight moves
			// Check all 8 possible moves
			for j := 0; j < 8; j++ {
				newX := current.X + dx[j]
				newY := current.Y + dy[j]
				newPos := Position{X: newX, Y: newY}

				if !visited[newPos] && isValidPosition(newPos) {
					visited[newPos] = true
					queue = append(queue, newPos)
				}
			}
			fmt.Println("Queue:", queue)
		}

		queue = queue[size:]
		steps++
	}

	return math.MaxInt32 // unreachable
}



func isValidPosition(pos Position) bool {
	// Assuming a standard chessboard size of 8x8
	return pos.X >= 0 && pos.X < 8 && pos.Y >= 0 && pos.Y < 8
}


