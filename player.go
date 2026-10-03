package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	playerLives               int
	posX, posY                int
	playerMoveSpeed           int
	image                     *ebiten.Image
	playerWidth, playerHeight int
}

func (p *Player) Move() {
	var nextX, nextY int
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		nextX = p.posX - p.playerMoveSpeed
		if nextX >= 0 {
			p.posX = nextX
		} else {
			p.posX = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		nextX = p.posX + p.playerMoveSpeed
		if nextX+p.playerWidth <= screenWidth {
			p.posX = nextX
		} else {
			p.posX = screenWidth - p.playerWidth
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		nextY = p.posY - p.playerMoveSpeed
		if nextY >= 0 {
			p.posY = nextY
		} else {
			p.posY = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		nextY = p.posY + p.playerMoveSpeed
		if nextY+p.playerHeight <= screenHeight {
			p.posY = nextY
		} else {
			p.posY = screenHeight - p.playerHeight
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	if p.playerLives == 0 {
		return
	}
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Scale(0.025, 0.025)

	playerOp.GeoM.Translate(float64(p.posX), float64(p.posY))
	screen.DrawImage(p.image, playerOp)
}

func (p *Player) ResetPos() {
	p.posX, p.posY = playerOriginX, playerOriginY
}
