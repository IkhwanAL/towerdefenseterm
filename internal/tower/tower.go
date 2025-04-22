package tower

import "github.com/gdamore/tcell/v2"

var TowerLocation = [][]int{
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

func GenerateTowerPlaceholder(
	towerLocation [][]int,
	screen tcell.Screen,
) {
	for i := range towerLocation {
		locationToPlace := towerLocation[i]

		screen.SetContent(locationToPlace[1]-1, locationToPlace[0]-1, ' ', nil, tcell.StyleDefault)
		screen.SetContent(locationToPlace[1], locationToPlace[0]-1, ' ', nil, tcell.StyleDefault)
		screen.SetContent(locationToPlace[1]+1, locationToPlace[0]-1, ' ', nil, tcell.StyleDefault)

		screen.SetContent(locationToPlace[1]-1, locationToPlace[0], ' ', nil, tcell.StyleDefault)
		screen.SetContent(locationToPlace[1], locationToPlace[0], ' ', nil, tcell.StyleDefault)
		screen.SetContent(locationToPlace[1]+1, locationToPlace[0], ' ', nil, tcell.StyleDefault)

		screen.SetContent(locationToPlace[1]-1, locationToPlace[0]+1, ' ', nil, tcell.StyleDefault)
		screen.SetContent(locationToPlace[1], locationToPlace[0]+1, ' ', nil, tcell.StyleDefault)
		screen.SetContent(locationToPlace[1]+1, locationToPlace[0]+1, ' ', nil, tcell.StyleDefault)
	}
}

func AllowedToPlaceTower(x, y int, towerLocation [][]int) (int, int) {
	locationAccepted := false

	for i := range towerLocation {
		location := towerLocation[i]

		if location[0]-1 == y || location[0]+1 == y {
			return -1, -1
		}

		if location[1]-1 == x || location[1]+1 == x {
			return -1, -1
		}

		if location[0] == y && location[1] == x {
			locationAccepted = true
			break
		}
	}

	if !locationAccepted {
		return -1, -1
	}

	return x, y

}

func PlaceATower(screen tcell.Screen, x, y int) {
	screen.SetContent(x-1, y-1, '╭', nil, tcell.StyleDefault)
	screen.SetContent(x, y-1, '-', nil, tcell.StyleDefault)
	screen.SetContent(x+1, y-1, '╮', nil, tcell.StyleDefault)

	screen.SetContent(x-1, y, '|', nil, tcell.StyleDefault)
	screen.SetContent(x, y, '*', nil, tcell.StyleDefault)
	screen.SetContent(x+1, y, '|', nil, tcell.StyleDefault)

	screen.SetContent(x-1, y+1, '╰', nil, tcell.StyleDefault)
	screen.SetContent(x, y+1, '-', nil, tcell.StyleDefault)
	screen.SetContent(x+1, y+1, '╯', nil, tcell.StyleDefault)
}
