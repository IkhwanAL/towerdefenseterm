package tower

import (
	"log"
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

func AllowedToPlaceTower(x, y int, towerLocation [][]int) (bool, []int) {
	locationAccepted := false

	locationPoint := make([]int, 2)

	for i := range towerLocation {
		location := towerLocation[i]

		yAllowed := y == location[0]
		xAllowed := x == location[1] || x == location[1]-1 || x == location[1]+1

		if xAllowed && yAllowed {
			locationAccepted = true
			locationPoint = location
			break
		}
	}

	return locationAccepted, locationPoint
}

func CheckForScreenBeforePlaceTower(x, y int, screen tcell.Screen) bool {
	leftContent, _, _, _ := screen.GetContent(x-1, y)
	centerContent, _, _, _ := screen.GetContent(x, y)
	rightContent, _, _, _ := screen.GetContent(x+1, y)

	if leftContent != ' ' {
		return false
	}

	if centerContent != ' ' {
		return false
	}

	if rightContent != ' ' {
		return false
	}

	return true
}

func (tower *Tower) Attack() int {
	tower.LastTimeAttack = time.Now()
	log.Printf("ATK %v", tower.LastTimeAttack)
	return 1
}

func (tower *Tower) CanAttackNow() bool {
	log.Printf("Attack Since %v", tower.LastTimeAttack)
	return time.Since(tower.LastTimeAttack) > tower.AttackSpeed
}

func PlaceATower(screen tcell.Screen, x, y int, attackSpeed time.Duration) *Tower {
	screen.SetContent(x-1, y, '|', nil, tcell.StyleDefault.Foreground(tcell.ColorLightSkyBlue))
	screen.SetContent(x, y, '*', nil, tcell.StyleDefault.Foreground(tcell.ColorLightSkyBlue))
	screen.SetContent(x+1, y, '|', nil, tcell.StyleDefault.Foreground(tcell.ColorLightSkyBlue))
	return &Tower{
		W:           x,
		H:           y,
		LOS:         4,
		AttackSpeed: attackSpeed,
	}
}

func euclideanFormula(px, py, qx, qy float64) int {
	x := math.Pow(qx-px, 2)
	y := math.Pow(qy-py, 2)

	return int(math.Sqrt(x + y))
}
