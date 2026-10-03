package main

import "github.com/hajimehoshi/ebiten/v2"

type Key int

const (
	KeyEnter Key = Key(ebiten.KeyEnter)
	KeyLeft  Key = Key(ebiten.KeyArrowLeft)
	KeyRight Key = Key(ebiten.KeyArrowRight)
	KeyUp    Key = Key(ebiten.KeyArrowUp)
	KeyDown  Key = Key(ebiten.KeyArrowDown)
	KeySpace Key = Key(ebiten.KeySpace)
)
