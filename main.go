package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	g.stateManager.HandleStateTransition()
	g.state = g.stateManager.currentState
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	renderer := &ScreenRenderer{}
	currentScreen := renderer.ScreenInitialize(g.state)
	screen.Clear()
	switch currentScreen {
	case StartScreen:
		// draw menu
	case PlayScreen:
		// draw game
	case EndScreen:
		// draw game over
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{
		state: StateMenu,
		lives: playerLives,
	}
	g.stateManager = &StateManager{
		currentState: StateMenu,
		game: g,
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Space Invaders")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}