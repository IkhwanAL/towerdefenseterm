package main

import "fmt"

const height = 25
const width = 120

var towerLocation = [][]int{
	{(25 / 2) - 4, 100},
	{(25 / 2) - 4, 90},
	{(25 / 2) - 4, 80},
	{(25 / 2) - 4, 50},
	{(25 / 2) - 4, 30},

	{(25 / 2) + 4, 30},
	{(25 / 2) + 4, 50},
	{(25 / 2) + 4, 80},
	{(25 / 2) + 4, 90},
	{(25 / 2) + 4, 100},
}

func generateTowerPlaceholder(towerLocation [][]int, gameMap *[height][width]string) {
	for i := range towerLocation {
		locationToPlace := towerLocation[i]

		gameMap[locationToPlace[0]-1][locationToPlace[1]-1] = " "
		gameMap[locationToPlace[0]-1][locationToPlace[1]] = " "
		gameMap[locationToPlace[0]-1][locationToPlace[1]+1] = " "
		gameMap[locationToPlace[0]][locationToPlace[1]-1] = " "
		gameMap[locationToPlace[0]][locationToPlace[1]+1] = " "
		gameMap[locationToPlace[0]][locationToPlace[1]] = "*"
		gameMap[locationToPlace[0]+1][locationToPlace[1]-1] = " "
		gameMap[locationToPlace[0]+1][locationToPlace[1]] = " "
		gameMap[locationToPlace[0]+1][locationToPlace[1]+1] = " "
	}
}

func main() {
	var gameMap [height][width]string

	// Generate Road
	for h := range height {
		for w := range width {
			leftJunction := height / 2
			road := (height / 2) - 1
			rightJunction := (height / 2) + 1

			if h == leftJunction || h == road || h == rightJunction {
				gameMap[h][w] = " "
			} else {
				gameMap[h][w] = "#"
			}
		}
	}

	generateTowerPlaceholder(towerLocation, &gameMap)

	// Set Fake Dimension
	// To Handle Out Of Index Array
	for h := range gameMap {
		for w := range gameMap[h] {
			marker := gameMap[h][w]

			fmt.Print(marker)

		}
		fmt.Print("\n")
	}
}
