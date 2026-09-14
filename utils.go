package main
import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (sm *StateManager) HandleStateTransition() {
	switch sm.currentState {
	case StateMenu:
		sm.currentState = StatePlaying
	case StatePlaying: 
		if sm.game.lives <= 0 {
			sm.currentState = StateGameOver
		}
	case StateGameOver: 
		if sm.game.lives > 0 {
			sm.currentState = StatePlaying
		}
	}
}

func (sc *ScreenRenderer) ScreenInitialize(state GameState) GameScreen {
	switch state {
	case StateMenu:
		return StartScreen
	case StatePlaying:
		return PlayScreen
	case StateGameOver:
		return EndScreen
	}
}

func (sc *ScreenRenderer) Render(ebScreen *ebiten.Image) {

}