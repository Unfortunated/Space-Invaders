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
		if g.debugTicks >= 60 {
			// log.Printf("state=%v playerPos=(%v,%v)", g.state, g.player.posX, g.player.posY)
			// log.Printf("state=%v bulletPos=(%v,%v) bulletActive=%v", g.state, g.bullets.posX, g.bullets.posY, g.bullets.isActive)
			g.debugTicks = 0
		}
		g.debugTicks++
		g.player.Move()
		activeBullets := 0
		for i := range g.bullets {
			g.bullets[i].Update()
			if g.bullets[i].isActive {
				g.bullets[activeBullets] = g.bullets[i]
				activeBullets += 1
			}
		}
		g.bullets = g.bullets[:activeBullets]
		var x, y, fired = g.player.Shoot()
		if fired && len(g.bullets) < 4 {
			g.bullets = append(g.bullets, Bullet{x, y, bulletSpeed, g.bulletImage, fired, bulletWidth})
		}
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
		for i := range g.bullets {
			if g.bullets[i].isActive {
				g.bullets[i].Draw(screen)
			}
		}
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

	bulletImage, _, bulletErr := ebitenutil.NewImageFromFile("assets/bullet.png")
	if bulletErr != nil {
		log.Fatal(bulletErr)
	}

	g := &Game{
		player:      Player{3, playerOriginX, playerOriginY, 3, spaceShip, 24, 48},
		state:       StateMenu,
		lives:       playerLives,
		bullets:     []Bullet{},
		bulletImage: bulletImage,
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
