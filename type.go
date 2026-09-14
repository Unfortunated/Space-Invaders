package main

type Game struct{
	state GameState
	score int
	lives int
	level int
	stateManager *StateManager
}

type StateManager struct {
	currentState GameState
	game *Game
}

type ScreenRenderer struct {
	screen GameScreen
}

type GameState int
const (
	StateMenu GameState = iota
	StatePlaying
	StateGameOver
)

type GameScreen int
const (
	StartScreen GameScreen = iota
	PlayScreen
	EndScreen
)