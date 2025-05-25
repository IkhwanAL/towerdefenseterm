package generator

import (
	"fmt"
	"math/rand"
	"testing"
)

var height = 100
var width = 24

func IsValidRoad(grid [][]rune, roadMap [][]int) error {

	for _, point := range roadMap {
		y, x := point[0], point[1]

		if y < 0 && y > height || x < 0 && x > width {
			return fmt.Errorf("(%d, %d): is out of bound", x, y)
		}

		if grid[y][x] != road {
			return fmt.Errorf("grid mismatch at road point (%d, %d): expected ' ', got '%c'", x, y, grid[y][x])
		}
	}

	return nil
}

func FuzzGenerateRoad(f *testing.F) {
	f.Add(int64(1748163368027334600))

	f.Fuzz(func(t *testing.T, seed int64) {
		rand := rand.New(rand.NewSource(seed))

		grid, roadMap := Road(rand, height, width)

		if err := IsValidRoad(grid, roadMap); err != nil {
			t.Fatalf("invalid output with seed %d: %v", seed, err)
		}
	})
}
