package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gdamore/tcell/v2"
)

const height = 25
const width = 120

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
	screen tcell.Screen,
) {
	for i := range towerLocation {
		locationToPlace := towerLocation[i]

		screen.SetContent(locationToPlace[1]-1, locationToPlace[0]-1, ' ', nil, tcell.StyleDefault)
		screen.SetContent(locationToPlace[1], locationToPlace[0]-1, ' ', nil, tcell.StyleDefault)
		screen.SetContent(locationToPlace[1]+1, locationToPlace[0]-1, ' ', nil, tcell.StyleDefault)

		screen.SetContent(locationToPlace[1]-1, locationToPlace[0], ' ', nil, tcell.StyleDefault)
		screen.SetContent(locationToPlace[1], locationToPlace[0], '*', nil, tcell.StyleDefault)
		screen.SetContent(locationToPlace[1]+1, locationToPlace[0], ' ', nil, tcell.StyleDefault)

		screen.SetContent(locationToPlace[1]-1, locationToPlace[0]+1, ' ', nil, tcell.StyleDefault)
		screen.SetContent(locationToPlace[1], locationToPlace[0]+1, ' ', nil, tcell.StyleDefault)
		screen.SetContent(locationToPlace[1]+1, locationToPlace[0]+1, ' ', nil, tcell.StyleDefault)
	}
}

func interrupt(screen tcell.Screen, notify chan os.Signal) {
	signal.Notify(notify, os.Interrupt)

	go func() {
		<-notify // Receive
		screen.Fini()
		os.Exit(0)
	}()
}
func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = screen.Init(); err != nil {
		log.Fatal(err.Error())
	}

	var notifyChan chan os.Signal = make(chan os.Signal, 1)

	interrupt(screen, notifyChan)

	screen.Clear()

	// Generate Road
	for h := range height {
		for w := range width {
			leftJunction := height / 2
			road := (height / 2) - 1
			rightJunction := (height / 2) + 1

			if h == leftJunction || h == road || h == rightJunction {
				screen.SetContent(w, h, ' ', nil, tcell.StyleDefault)
			} else {
				screen.SetContent(w, h, '#', nil, tcell.StyleDefault)
			}
		}
	}

	generateTowerPlaceholder(towerLocation, screen)

	enemies := GenerateEnemy()

	for _, enemy := range enemies {
		screen.SetContent(enemy.W, enemy.H, enemy.Type, nil, tcell.StyleDefault)
	}

	screen.Show()
	frameTime := time.NewTicker(500 * time.Millisecond)
	defer frameTime.Stop()

	// TODO Need To Keep Track The Total Enemy And Their State
	for {
		select {
		case <-frameTime.C:

			for _, enemy := range enemies {
				enemy.GoLeft()
				screen.SetContent(enemy.W, enemy.H, enemy.Type, nil, tcell.StyleDefault)
				screen.SetContent(enemy.W-1, enemy.H, ' ', nil, tcell.StyleDefault)
			}

			screen.Show()
		default:
			if screen.HasPendingEvent() {
				ev := screen.PollEvent()
				switch ev := ev.(type) {
				case *tcell.EventKey:
					if ev.Key() == tcell.KeyEscape || ev.Rune() == 'q' {
						screen.Fini()
						os.Exit(0)
					}
				}
			}
		}
	}
}
