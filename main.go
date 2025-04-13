package main

import (
	"fmt"
)

const height = 25
const width = 120

const enemy = "█"

var towerLocation = [][]int{
	{(25 / 2) - 4, 100, 4},
	{(25 / 2) - 4, 90, 4},
	{(25 / 2) - 4, 80, 4},
	{(25 / 2) - 4, 50, 4},
	{(25 / 2) - 4, 30, 4},

	{(25 / 2) + 4, 30, 4},
	{(25 / 2) + 4, 50, 4},
	{(25 / 2) + 4, 80, 4},
	{(25 / 2) + 4, 90, 4},
	{(25 / 2) + 4, 100, 4},
}

func generateTowerPlaceholder(
	towerLocation [][]int,
	gameMap *[height][width]string,
) {
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

func isInsideTowerLocation(h, w int, towerLocation [][]int) bool {
	for i := range towerLocation {
		location := towerLocation[i]

		hLocation := location[0]
		wLocation := location[1]

		if h == hLocation-1 {
			return true
		}

		if h == hLocation {
			return true
		}

		if h == hLocation+1 {
			return true
		}

		if w == wLocation-1 {
			return true
		}

		if w == wLocation {
			return true
		}

		if w == wLocation+1 {
			return true
		}
	}

	return false
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

	//TODO Update Every frame

	totalEnemy := 0
	const maxTotalEnemy = 1

	// for {
	for h := range gameMap {
		for w := range gameMap[h] {
			marker := gameMap[h][w]

			inTowerLocation := isInsideTowerLocation(h, w, towerLocation)

			// How to Render Enemy
			if marker == " " && totalEnemy < maxTotalEnemy && !inTowerLocation {
				fmt.Print(enemy)
				totalEnemy += 1
				continue
			}

			fmt.Print(marker)

		}
		fmt.Print("\n")
	}

	// }
}
