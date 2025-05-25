package enemy

import (
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
)

const GRUNT = '○'

type Enemy struct {
	H            int
	W            int
	Type         rune
	HP           int
	Interval     time.Duration
	LastMoved    time.Time
	Color        []int
	LastTimeHit  time.Time
	TickFlashing time.Duration
	RoadState    int
	Alive        bool
}

func (enemy *Enemy) TakeDamage(amount int) {
	enemy.HP -= amount
	enemy.LastTimeHit = time.Now()
	log.Printf("Hit At %v: ", enemy.LastTimeHit)

	if enemy.HP <= 0 {
		enemy.Alive = false
	}
}

func (enemy *Enemy) MustFlashing() bool {
	log.Printf("Hit Since %v: ", enemy.LastTimeHit)
	return time.Since(enemy.LastTimeHit) < enemy.TickFlashing
}

func (enemy *Enemy) Draw(screen tcell.Screen) {
	color := tcell.NewRGBColor(
		int32(enemy.Color[0]),
		int32(enemy.Color[1]),
		int32(enemy.Color[2]),
	)

	enemyType := enemy.Type

	if !enemy.Alive {
		enemyType = ' '
	}

	if enemy.MustFlashing() {
		color = tcell.ColorRed
	}

	screen.SetContent(enemy.W, enemy.H, enemyType, nil, tcell.StyleDefault.Foreground(color))
}

func GenerateEnemy(baseInterval time.Duration, height int, flashTick time.Duration, hStart, wStart int) []*Enemy {
	now := time.Now()

	return []*Enemy{
		{
			H:            hStart,
			W:            wStart,
			Type:         GRUNT,
			HP:           2,
			Interval:     baseInterval * 4,
			LastMoved:    now,
			Color:        []int{0, 0, 255}, // Blue
			TickFlashing: flashTick,
			RoadState:    0,
			Alive:        true,
		},
		{
			H:            hStart,
			W:            wStart - 2,
			Type:         GRUNT,
			HP:           3,
			Interval:     baseInterval * 5,
			LastMoved:    now.Add(baseInterval * 5),
			Color:        []int{0, 0, 255}, // Blue
			TickFlashing: flashTick,
			RoadState:    0,
			Alive:        true,
		},
		{
			H:            hStart,
			W:            wStart - 2,
			Type:         GRUNT,
			HP:           3,
			Interval:     baseInterval * 6,
			LastMoved:    now.Add(baseInterval * 15),
			Color:        []int{0, 0, 255}, // Blue
			TickFlashing: flashTick,
			RoadState:    0,
			Alive:        true,
		},
	}
}
