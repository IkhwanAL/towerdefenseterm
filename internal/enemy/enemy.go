package enemy

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

const GRUNT = '█'

type Enemy struct {
	H             int
	W             int
	Type          rune
	HP            int
	Interval      time.Duration
	LastMoved     time.Time
	Color         []int
	LastTimeHit   time.Time
	FlashDuration time.Duration
}

func (enemy *Enemy) GoLeft() {
	enemy.W -= 2
}

func (enemy *Enemy) GoTop() {
	enemy.H += 2
}

func (enemy *Enemy) GoBottom() {
	enemy.H -= 2
}

func (enemy *Enemy) GoRight() {
	enemy.W += 2
}

func (enemy *Enemy) TakeDamage(amount int) {
	enemy.HP -= amount
	enemy.LastTimeHit = time.Now()
	enemy.FlashDuration = 400 * time.Millisecond
}

func (enemy *Enemy) IsFlashing() bool {
	return time.Since(enemy.LastTimeHit) < enemy.FlashDuration
}

func (enemy *Enemy) Draw(screen tcell.Screen) {
	color := tcell.NewRGBColor(
		int32(enemy.Color[0]),
		int32(enemy.Color[1]),
		int32(enemy.Color[2]),
	)

	if enemy.IsFlashing() {
		color = tcell.ColorRed
	}

	screen.SetContent(enemy.W, enemy.H, enemy.Type, nil, tcell.StyleDefault.Foreground(color))
}

func GenerateEnemy(baseInterval time.Duration, height int) []*Enemy {
	now := time.Now()

	return []*Enemy{
		{
			H:         height / 2,
			W:         -2,
			Type:      GRUNT,
			HP:        3,
			Interval:  baseInterval,
			LastMoved: now,
			Color:     []int{0, 0, 255}, // Blue
		},
		// {
		// 	H:         height / 2,
		// 	W:         -2,
		// 	Type:      GRUNT,
		// 	HP:        3,
		// 	Interval:  baseInterval * 2,
		// 	LastMoved: now.Add(baseInterval * 5),
		// 	Color:     []int{0, 0, 255}, // Blue
		// },
		// {
		// 	H:         height / 2,
		// 	W:         -2,
		// 	Type:      GRUNT,
		// 	HP:        3,
		// 	Interval:  baseInterval * 4,
		// 	LastMoved: now.Add(baseInterval * 15),
		// 	Color:     []int{0, 0, 255}, // Blue
		// },
	}
}
