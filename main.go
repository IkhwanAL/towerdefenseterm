package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/IkhwanAL/towerdefenseterm/internal/enemy"
	"github.com/IkhwanAL/towerdefenseterm/internal/tower"
	"github.com/gdamore/tcell/v2"
)

const height = 25
const width = 120

const losRad = 5

func interrupt(screen tcell.Screen, notify chan os.Signal) {
	signal.Notify(notify, os.Interrupt)

	go func() {
		<-notify // Receive
		screen.Fini()
		os.Exit(0)
	}()
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

	tower.GenerateTowerPlaceholder(tower.TowerLocation, screen)

	tick := 300 * time.Millisecond

	enemies := enemy.GenerateEnemy(tick, height)

	for _, enemy := range enemies {
		screen.SetContent(enemy.W, enemy.H, ' ', nil, tcell.StyleDefault)
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

	var availableTower []*tower.Tower

	// TODO Able To Shoot And Unit Take Damage
	// TODO Able to Make Damage Tick With Red Color If it's get Hit
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
					x, y := ev.Position()

					x, y = tower.AllowedToPlaceTower(x, y, tower.TowerLocation)

					if x == -1 && y == -1 {
						break
					}

					createdTower := tower.PlaceATower(screen, x, y, 2000)

					availableTower = append(availableTower, createdTower)

					log.Print("I Place Tower")
				}
			}

			screen.Show()
		case <-frameTime.C:

			now := time.Now()

			var enemyMoved []*enemy.Enemy

			for index := range enemies {
				enemy := enemies[index]

				log.Printf("HP %d", enemy.HP)

				lastMoved := now.Sub(enemy.LastMoved)

				if lastMoved >= enemy.Interval {
					screen.SetContent(enemy.W, enemy.H, ' ', nil, tcell.StyleDefault) // Removing Track

					enemy.GoRight()
					enemy.LastMoved = now
					enemy.Draw(screen)
					enemyMoved = append(enemyMoved, enemy)
				}
			}

			for _, watchTower := range availableTower {
				// watchTower := availableTower[i]

				for _, target := range enemyMoved {
					//target := enemyMoved[j]

					isInArea := watchTower.UnitCloseToTower(
						float64(target.W),
						float64(target.H),
						float64(watchTower.W),
						float64(watchTower.H),
					)

					if isInArea && watchTower.CanAttackNow() {
						target.TakeDamage(watchTower.Attack())
						target.Draw(screen)
					}

				}
			}

			screen.Show()
		}
	}
}
