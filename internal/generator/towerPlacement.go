package generator

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

var road rune = ' '

func TowerPlacement(width, height, maxTower int, stateMap tcell.Screen) [][]int {
	var heightLocations []int
	var widthLocations []int

	for h := range height {
		for w := range width {
			pixelState, _, _, _ := stateMap.GetContent(w, h)

			if pixelState != road {
				continue
			}

			// the minus and plus three are
			// the padding between middle point and the road
			heightLocations = append(heightLocations, h-3)
			heightLocations = append(heightLocations, h+3)
			widthLocations = append(widthLocations, w)
		}
	}

	totalTower := 0

	var towerPlacement [][]int

	for totalTower < maxTower {
		w := getRandomizePoint(widthLocations)
		h := getRandomizePoint(heightLocations)

		stateMap.SetContent(w, h, '$', nil, tcell.StyleDefault)

		towerPlacement = append(towerPlacement, []int{h, w})

		totalTower += 1
	}

	return towerPlacement
}

func getRandomizePoint(listLocation []int) int {
	return listLocation[rand.Intn(len(listLocation))]
}
