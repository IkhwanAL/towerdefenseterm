package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/IkhwanAL/towerdefenseterm/internal/enemy"
	"github.com/IkhwanAL/towerdefenseterm/internal/generator"
	"github.com/IkhwanAL/towerdefenseterm/internal/tower"
	"github.com/gdamore/tcell/v2"
)

const height = 25
const width = 180

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

	grid := generator.Road(width, height)

	// Generate Road
	for h, row := range grid {
		for w, ch := range row {
			screen.SetContent(w, h, ch, nil, tcell.StyleDefault)
		}
	}

	towerLocation := generator.TowerPlacement(width, height, 12, screen)

	tick := 100 * time.Millisecond

	enemies := enemy.GenerateEnemy(tick, height, tick*4)

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

	//TODO Random Generatar Road
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

					accepted, locationPoint := tower.AllowedToPlaceTower(x, y, towerLocation)

					if !accepted {
						log.Printf("%d, %d :: Failed Location Pick", x, y)
						break
					}

					if !tower.CheckForScreenBeforePlaceTower(locationPoint[1], locationPoint[0], screen) {
						log.Printf("%d, %d :: Failed Location Pick", x, y)
						break
					}

					log.Printf("%v :: Accepted Location Point", locationPoint)

					createdTower := tower.PlaceATower(screen, locationPoint[1], locationPoint[0], tick*7)

					availableTower = append(availableTower, createdTower)

					log.Print("I Place Tower")
				}
			}

			screen.Show()
		case <-frameTime.C:

			now := time.Now()

			var enemyMoved []*enemy.Enemy
			var restOfEnemy []*enemy.Enemy

			for index := range enemies {
				enemy := enemies[index]

				// log.Printf("HP %d", enemy.HP)

				lastMoved := now.Sub(enemy.LastMoved)

				if lastMoved >= enemy.Interval {
					screen.SetContent(enemy.W, enemy.H, ' ', nil, tcell.StyleDefault) // Removing Track

					enemy.GoRight()
					enemy.LastMoved = now
					enemyMoved = append(enemyMoved, enemy)
					enemy.Draw(screen)
				} else {
					restOfEnemy = append(restOfEnemy, enemy)
				}
			}

			var stillAliveEnemies []*enemy.Enemy

			if len(availableTower) == 0 {
				stillAliveEnemies = enemyMoved
			}

			for _, watchTower := range availableTower {
				for _, target := range enemyMoved {
					isInArea := watchTower.UnitCloseToTower(
						float64(target.W),
						float64(target.H),
						float64(watchTower.W),
						float64(watchTower.H),
					)

					if isInArea && watchTower.CanAttackNow() {
						target.TakeDamage(watchTower.Attack())
					}

					target.Draw(screen)
					if target.HP > 0 {
						stillAliveEnemies = append(stillAliveEnemies, target)
					}
				}
			}
			screen.Show()

			restOfEnemy = append(restOfEnemy, stillAliveEnemies...)
			enemies = restOfEnemy
			// log.Printf("Not Moved %v", restOfEnemy)
			// log.Printf("Just Moved %v", enemyMoved)
			// log.Printf("Moved And Still Alive %v", stillAliveEnemies)

		}
	}
}
