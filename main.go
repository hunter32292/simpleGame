package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hunter32292/simpleGame/pkg/game"
)

var simpleGame game.Game

func main() {
	log.Println("Starting simple game ...")
	setupGame()
	gameLoop()
	cleanUp()
	log.Println("Game exit")
}

// Setup the Game
func setupGame() {
	p := &game.Player{
		Money:     200,
		Items:     []game.Items{},
		Health:    100,
		LocationX: 0,
		LocationY: 0,
	}
	simpleGame = game.Game{
		Clock:  0,
		Player: p,
	}
	// TODO: Add intro Screen here
}

// Start the game
func gameLoop() {
	var win = false
	// Infinite Game loop with check
	for {
		gamePrompt()
		// TODO: draw game from window
		// screen.DrawGame(simpleGame)

		event := game.GenerateEvent()
		playEvent(event, &simpleGame)

		if simpleGame.Player.Health <= 0 {
			log.Println("You've died.")
			break
		}
		// If we've won the game
		if win != false {
			log.Println("You've won the game!")
			break
		}
		simpleGame.Clock++
	}
	return
}

// Clean Up the game
func cleanUp() {

}

func playEvent(event game.Event, sg *game.Game) {
	var answer string
	fmt.Println(event.Name)
	fmt.Println(event.Question)
	for _, o := range event.Options {
		fmt.Println(o)
	}
	fmt.Scanln(&answer)

	if !game.ContainsOption(event.Options, answer) {
		return
	}
	if answer == event.Answer {
		sg.Player.Money += event.Value
		sg.Player.Health += event.Damage
		fmt.Println(event.Outcome(answer))
		time.Sleep(time.Second * 2)
	} else {
		fmt.Println(event.Outcome(answer))
		time.Sleep(time.Second * 2)
	}
}

// Setup Game for playing
func gamePrompt() {
	fmt.Printf("\x1b[H\x1B[2J") // Clear screen for the current event
	fmt.Print("Time: ")
	fmt.Println(simpleGame.Clock)
	fmt.Println("--- Player ---")
	fmt.Println(simpleGame.Player)
}
