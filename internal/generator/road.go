package generator

import (
	"log"
	"math/rand"
)

var land rune = '#'

func Road(rand *rand.Rand, width, height int) ([][]rune, [][]int) {
	grid := make([][]rune, height)

	for h := range grid {
		grid[h] = make([]rune, width)
		for w := range grid[h] {
			grid[h][w] = land
		}
	}

	y := rand.Intn(height)
	grid[y][0] = road

	log.Printf("Start Height (y) :: %d", y)
	var prevCurves rune

	var pointLocationOfRoad [][]int

	for x := 1; x < width; x++ {
		var curvesUsed rune

		// The Shifting Part is Where Road Change Direction (for now up or down)
		// The bigger the max shifting number the seldom it to change direction
		shift := rand.Intn(7)
		log.Printf("Shift :: %d", shift)
		if shift == 0 && y > 6 {
			y-- // go up
			grid[y][x-1] = '/'
			grid[y+1][x] = '/'
			curvesUsed = '/'
		} else if shift == 1 && y < height-6 {
			y++ // go down
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
		pointLocationOfRoad = append(pointLocationOfRoad, []int{y, x})
		prevCurves = curvesUsed
	}

	return grid, pointLocationOfRoad
}
