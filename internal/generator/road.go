package generator

import (
	"math/rand"
)

var land rune = '#'

func Road(width, height int) [][]rune {
	grid := make([][]rune, height)

	for h := range grid {
		grid[h] = make([]rune, width)
		for w := range grid[h] {
			grid[h][w] = land
		}
	}

	y := rand.Intn(height)
	// prevShift := 99

	for x := range width {
		// grid[randomHeight][w] = road
		// grid[randomHeight+1][w] = road

		shift := rand.Intn(6)
		if shift == 0 && y > 6 {
			y--
			grid[y][x-1] = '/'
			grid[y+1][x] = '/'
		} else if shift == 1 && y < height-6 {
			y++
			grid[y-1][x] = '\\'
			grid[y][x-1] = '\\'
		}

		grid[y][x] = road
	}

	return grid
}
