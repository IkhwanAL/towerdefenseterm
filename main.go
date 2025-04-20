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
		screen.SetContent(locationToPlace[1], locationToPlace[0], ' ', nil, tcell.StyleDefault)
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

func placeTower(screen tcell.Screen, ev *tcell.EventMouse, towerLocation [][]int) {
	x, y := ev.Position()

	locationAccepted := false

	for i := range towerLocation {
		location := towerLocation[i]

		if location[0]-1 == y || location[0]+1 == y {
			return
		}

		if location[1]-1 == x || location[1]+1 == x {
			return
		}

		if location[0] == y && location[1] == x {
			locationAccepted = true
			break
		}
	}

	if !locationAccepted {
		return
	}

	screen.SetContent(x-1, y-1, '*', nil, tcell.StyleDefault)
	screen.SetContent(x, y-1, '*', nil, tcell.StyleDefault)
	screen.SetContent(x+1, y-1, '*', nil, tcell.StyleDefault)

	screen.SetContent(x-1, y, '*', nil, tcell.StyleDefault)
	screen.SetContent(x, y, '*', nil, tcell.StyleDefault)
	screen.SetContent(x+1, y, '*', nil, tcell.StyleDefault)

	screen.SetContent(x-1, y+1, '*', nil, tcell.StyleDefault)
	screen.SetContent(x, y+1, '*', nil, tcell.StyleDefault)
	screen.SetContent(x+1, y+1, '*', nil, tcell.StyleDefault)

}

func main() {
	logFile, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer logFile.Close()

	log.SetOutput(logFile)

	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatal(err.Error())
	}

	var notifyChan chan os.Signal = make(chan os.Signal, 1)

	interrupt(screen, notifyChan)

	if err = screen.Init(); err != nil {
		log.Fatal(err.Error())
	}

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

	tick := 500 * time.Millisecond

	enemies := GenerateEnemy(tick)

	for _, enemy := range enemies {
		screen.SetContent(enemy.W, enemy.H, enemy.Type, nil, tcell.StyleDefault)
	}

	screen.Show()
	screen.EnableMouse()

	frameTime := time.NewTicker(tick)
	defer frameTime.Stop()

	eventChan := make(chan tcell.Event, 1)

	go func() {
		for {
			ev := screen.PollEvent()
			eventChan <- ev
		}
	}()

	for {
		select {
		case ev := <-eventChan:
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Rune() == 'q' {
					screen.Fini()
					os.Exit(0)
					break
				}
			case *tcell.EventMouse:
				if ev.Buttons() == tcell.Button1 {
					placeTower(screen, ev, towerLocation)
				}
			}

			screen.Show()
		case <-frameTime.C:
			// TODO Still Can't Handle More than 2 Unit

			now := time.Now()

			for index := range enemies {
				enemy := &enemies[index]

				lastMoved := now.Sub(enemy.LastMoved)

				log.Printf("Unit: %d", index)
				log.Printf("Last Moved: %s ", lastMoved)
				log.Printf("Last Moved (Time): %s", enemy.LastMoved)
				log.Printf("interval: %s", enemy.Interval)
				log.Printf("Tick (Time): %s", now)

				if lastMoved >= enemy.Interval {
					screen.SetContent(enemy.W, enemy.H, ' ', nil, tcell.StyleDefault) // Removing Track

					enemy.GoRight()
					enemy.LastMoved = now
					color := tcell.NewRGBColor(
						int32(enemy.Color[0]),
						int32(enemy.Color[1]),
						int32(enemy.Color[2]),
					)
					screen.SetContent(enemy.W, enemy.H, enemy.Type, nil, tcell.StyleDefault.Foreground(color))
				}
				log.Print("\n")
			}

			screen.Show()
		}
	}
}
