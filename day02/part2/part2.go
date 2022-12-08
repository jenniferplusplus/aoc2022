package part2

import (
	"github.com/qmuntal/stateless"
	"strings"
)

const (
	rockState     = "A"
	paperState    = "B"
	scissorsState = "C"
	rockPlay      = 1
	paperPlay     = 2
	scissorsPlay  = 3
	loseMove      = "X"
	drawMove      = "Y"
	winMove       = "Z"
)

func Score(line string) int {
	strats := strings.Split(line, " ")
	game := gameState(strats[0])
	game.Fire(strats[1])

	score := 0
	switch strats[1] {
	case "X":
		score += 0
	case "Y":
		score += 3
	case "Z":
		score += 6
	}
	score += game.MustState().(int)

	return score
}

func gameState(shape string) *stateless.StateMachine {
	gameState := stateless.NewStateMachine(shape)
	gameState.Configure(rockState).
		Permit(winMove, paperPlay).
		Permit(drawMove, rockPlay).
		Permit(loseMove, scissorsPlay)
	gameState.Configure(paperState).
		Permit(winMove, scissorsPlay).
		Permit(drawMove, paperPlay).
		Permit(loseMove, rockPlay)
	gameState.Configure(scissorsState).
		Permit(winMove, rockPlay).
		Permit(drawMove, scissorsPlay).
		Permit(loseMove, paperPlay)

	return gameState
}
