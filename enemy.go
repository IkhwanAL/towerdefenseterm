package main

const GRUNT = '█'

type Enemy struct {
	H    int
	W    int
	Type rune
	HP   int
}

func (enemy *Enemy) GoLeft() {
	enemy.W += 2
}

func (enemy *Enemy) GoTop() {
	enemy.H += 1
}

func (enemy *Enemy) GoBottom() {
	enemy.H -= 1
}

func (enemy *Enemy) GoRight() {
	enemy.W -= 1
}

func GenerateEnemy() []*Enemy {
	return []*Enemy{
		{
			H:    height / 2,
			W:    0,
			Type: GRUNT,
			HP:   3,
		},
		{
			H:    height / 2,
			W:    0,
			Type: GRUNT,
			HP:   3,
		},
		{
			H:    height / 2,
			W:    0,
			Type: GRUNT,
			HP:   3,
		},
		{
			H:    height / 2,
			W:    0,
			Type: GRUNT,
			HP:   3,
		},
		{
			H:    height / 2,
			W:    0,
			Type: GRUNT,
			HP:   3,
		},
	}
}
