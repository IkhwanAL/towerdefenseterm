package generator

import "math/rand"

var land rune = '#'

func Road(width, height int) [][]rune {
	grid := make([][]rune, height)

	for h := range grid {
		grid[h] = make([]rune, width)
		for w := range grid[h] {
			grid[h][w] = land
		}
	}

	randomHeight := rand.Intn(height)
	for w := range width {
		grid[randomHeight][w] = road

		shift := rand.Intn(7)
		if shift == 0 && randomHeight > 6 {
			randomHeight--
			grid[randomHeight][w] = road
		} else if shift == 1 && randomHeight < height-6 {
			randomHeight++
			grid[randomHeight][w] = road
		}
	}

	return grid
}
