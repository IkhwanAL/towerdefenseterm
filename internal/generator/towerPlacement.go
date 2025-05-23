package generator

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

var road rune = ' '

func TowerPlacement(width, height, maxTower int, stateMap tcell.Screen) [][]int {
	var pixelLocation [][]int

	for h := range height {
		for w := range width {
			pixelState, _, _, _ := stateMap.GetContent(w, h)

			if pixelState != road {
				continue
			}

			pixelLocation = append(pixelLocation, []int{h, w})
		}
	}

	totalTower := 0

	var towerPlacement [][]int

	for totalTower < maxTower {
		dot := getRandomizePoint(pixelLocation)

		padding := 3

		negativeOrPositive := rand.Intn(2)

		if negativeOrPositive == 0 {
			padding = -padding
		}

		stateMap.SetContent(
			dot[1],
			dot[0]+padding,
			' ',
			nil,
			tcell.StyleDefault.Background(tcell.ColorAntiqueWhite),
		)

		stateMap.SetContent(
			dot[1]-1,
			dot[0]+padding,
			' ',
			nil,
			tcell.StyleDefault,
		)

		stateMap.SetContent(
			dot[1]+1,
			dot[0]+padding,
			' ',
			nil,
			tcell.StyleDefault,
		)

		towerPlacement = append(towerPlacement, []int{dot[0] + padding, dot[1]})

		totalTower += 1
	}

	return towerPlacement
}

func getRandomizePoint(listLocation [][]int) []int {
	return listLocation[rand.Intn(len(listLocation))]
}
