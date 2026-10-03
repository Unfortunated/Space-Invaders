package main

import "github.com/hajimehoshi/ebiten/v2"

type Bullet struct {
	posX, posY  int
	speed       int
	image       *ebiten.Image
	isActive    bool
	bulletWidth int
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	if !b.isActive {
		return
	}
	bulletOp := &ebiten.DrawImageOptions{}
	bulletOp.GeoM.Scale(0.01, 0.01)
	bulletOp.GeoM.Translate(float64(b.posX), float64(b.posY))
	screen.DrawImage(b.image, bulletOp)
}

func (b *Bullet) Update() {
	var nextY = b.posY - b.speed
	if nextY < 0 {
		b.isActive = false
		return
	} else {
		b.posY = nextY
	}
}
