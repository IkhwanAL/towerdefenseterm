package enemy

import "time"

const GRUNT = '█'

type Enemy struct {
	H         int
	W         int
	Type      rune
	HP        int
	Interval  time.Duration
	LastMoved time.Time
	Color     []int
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

func GenerateEnemy(baseInterval time.Duration, height int) []Enemy {
	now := time.Now()

	return []Enemy{
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
		// 	Color:     []int{51, 255, 51}, // Green
		// },
		// {
		// 	H:         height / 2,
		// 	W:         -2,
		// 	Type:      GRUNT,
		// 	HP:        3,
		// 	Interval:  baseInterval * 4,
		// 	LastMoved: now.Add(baseInterval * 15),
		// 	Color:     []int{255, 0, 0},
		// },
	}
}
