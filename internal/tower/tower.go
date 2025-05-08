package tower

import (
	"math"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Tower struct {
	W              int
	H              int
	LOS            int
	AttackSpeed    time.Duration
	LastTimeAttack time.Time
}

func (tower *Tower) UnitCloseToTower(px, py, qx, qy float64) bool {
	unitPosition := euclideanFormula(px, py, qx, qy)

	return unitPosition <= tower.LOS
}

// [H, W]
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

func (tower *Tower) Attack() int {
	tower.LastTimeAttack = time.Now()
	return 1
}

func (tower *Tower) CanAttackNow() bool {
	return true
}

func PlaceATower(screen tcell.Screen, x, y int, attackSpeed int) *Tower {
	screen.SetContent(x-1, y-1, '╭', nil, tcell.StyleDefault)
	screen.SetContent(x, y-1, '-', nil, tcell.StyleDefault)
	screen.SetContent(x+1, y-1, '╮', nil, tcell.StyleDefault)

	screen.SetContent(x-1, y, '|', nil, tcell.StyleDefault)
	screen.SetContent(x, y, '*', nil, tcell.StyleDefault)
	screen.SetContent(x+1, y, '|', nil, tcell.StyleDefault)

	screen.SetContent(x-1, y+1, '╰', nil, tcell.StyleDefault)
	screen.SetContent(x, y+1, '-', nil, tcell.StyleDefault)
	screen.SetContent(x+1, y+1, '╯', nil, tcell.StyleDefault)

	return &Tower{
		W:           x,
		H:           y,
		LOS:         4,
		AttackSpeed: time.Duration(attackSpeed * int(time.Millisecond)),
	}
}

func euclideanFormula(px, py, qx, qy float64) int {
	x := math.Pow(qx-px, 2)
	y := math.Pow(qy-py, 2)

	return int(math.Sqrt(x + y))
}
