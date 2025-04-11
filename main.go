package main

import "fmt"

const height = 25
const width = 120

func generateSquare(height, width, middle int) {

}

func main() {
	var excludeTowerLocation []int

	// Generate Road
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			leftJunction := height / 2
			road := (height / 2) - 1
			rightJunction := (height / 2) + 1

			if h == leftJunction || h == road || h == rightJunction {
				fmt.Print(" ")
				excludeTowerLocation = append(excludeTowerLocation, h)
			} else {
				fmt.Print("#")
			}
		}

		fmt.Print("\n")
	}
}
