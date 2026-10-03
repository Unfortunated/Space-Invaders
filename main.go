package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Update() error {
	g.stateManager.HandleStateTransition()
	g.state = g.stateManager.currentState

	if g.state == StatePlaying {
		if g.debugTicks >= 30 {
			log.Printf("state=%v playerPos=(%v,%v)", g.state, g.player.posX, g.player.posY)
			g.debugTicks = 0
		}
		g.debugTicks++
		g.player.Move()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	renderer := &ScreenRenderer{}
	currentScreen := renderer.ScreenInitialize(g.state)
	screen.Clear()
	switch currentScreen {
	case StartScreen:
		log.Println("Draw start screen")
	case PlayScreen:
		// game screen
		g.player.Draw(screen)
	case EndScreen:
		// draw game over
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	var spaceShip, _, err = ebitenutil.NewImageFromFile("assets/spaceship.png")
	if err != nil {
		log.Fatal(err)
	}

	g := &Game{
		player: Player{3, playerOriginX, playerOriginY, 3, spaceShip, 24, 48},
		state:  StateMenu,
		lives:  playerLives,
	}
	g.stateManager = &StateManager{
		currentState: StateMenu,
		game:         g,
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Space Invaders")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
