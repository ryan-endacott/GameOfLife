package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	grid := createGrid(10)

	fmt.Println("Starting Conway's Game of Life....")
	printGrid(grid)
	timer := time.Tick(time.Second)
	for now := range timer {
		updateGame(grid)
		fmt.Printf("Time is %d\n", now)
		printGrid(grid)
	}
}

// Performs one iteration of Conway's
// Game of Life on the grid
func updateGame(grid [][]int) {

	// Store during iteration because all steps happen at once
	cellsToLive := make([][2]int, 0)
	cellsToDie := make([][2]int, 0)

	// Determine who lives and who dies
	for x, row := range grid {
		for y, cell := range row {
			liveNeighborCount := 0
			for _, neighbor := range getNeighbors(grid, x, y) {
				if neighbor == 1 {
					liveNeighborCount += 1
				}
			}

			if cell == 1 { // if alive
				if liveNeighborCount < 2 { // underpop
					cellsToDie = append(cellsToDie, [2]int{x, y})
				}
				if liveNeighborCount > 3 { // overpop
					cellsToDie = append(cellsToDie, [2]int{x, y})
				}
			} else { // if dead
				if liveNeighborCount == 3 { // reproduction
					cellsToLive = append(cellsToLive, [2]int{x, y})
				}
			}
		}
	}

	// Give life and take it away
	for _, cell := range cellsToLive {
		grid[cell[0]][cell[1]] = 1
	}
	for _, cell := range cellsToDie {
		grid[cell[0]][cell[1]] = 0
	}

}

// Returns a slice of all neighbors to given cell
func getNeighbors(grid [][]int, x int, y int) (neighbors []int) {
	size := len(grid)
	neighbors = make([]int, 0, 8)

	// Add all in bounds neighbors
	if x-1 >= 0 && y-1 >= 0 {
		neighbors = append(neighbors, grid[x-1][y-1])
	}
	if x-1 >= 0 {
		neighbors = append(neighbors, grid[x-1][y])
	}
	if y-1 >= 0 {
		neighbors = append(neighbors, grid[x][y-1])
	}
	if x+1 < size && y+1 < size {
		neighbors = append(neighbors, grid[x+1][y+1])
	}
	if x+1 < size {
		neighbors = append(neighbors, grid[x+1][y])
	}

	if y+1 < size {
		neighbors = append(neighbors, grid[x][y+1])
	}
	if x+1 < size && y-1 >= 0 {
		neighbors = append(neighbors, grid[x+1][y-1])
	}
	if x-1 >= 0 && y+1 < size {
		neighbors = append(neighbors, grid[x-1][y+1])
	}

	return neighbors
}

// 0 is dead, 1 is alive
func createGrid(size int) [][]int {
	grid := make([][]int, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}

	// Randomly populate the grid
	for i := 0; i < size*3; i++ {
		x := rand.Intn(size)
		y := rand.Intn(size)
		grid[x][y] = 1
	}

	return grid
}

func printGrid(grid [][]int) {
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n") // spacing
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%d ", cell)
		}
		fmt.Printf("\n")
	}
}
