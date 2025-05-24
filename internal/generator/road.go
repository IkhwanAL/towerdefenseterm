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
	var prevCurves rune

	for x := range width {
		var curvesUsed rune

		shift := rand.Intn(7)
		if shift == 0 && y > 6 {
			y--
			grid[y][x-1] = '/'
			grid[y+1][x] = '/'
			curvesUsed = '/'
		} else if shift == 1 && y < height-6 {
			y++
			grid[y-1][x] = '\\'
			grid[y][x-1] = '\\'
			curvesUsed = '\\'
		}

		if prevCurves == '\\' && curvesUsed == '/' {
			grid[y][x-1] = 'V'
		}

		if prevCurves == '/' && curvesUsed == '\\' {
			grid[y][x-1] = 'Λ'
		}

		grid[y][x] = road
		prevCurves = curvesUsed
	}

	return grid
}
