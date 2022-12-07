package part1

import (
	"context"
	"github.com/qmuntal/stateless"
	"reflect"
	"strings"
)

const (
	rockState     = "A"
	paperState    = "B"
	scissorsState = "C"
	winState      = "Win"
	loseState     = "Lose"
	drawState     = "Draw"
	rockMove      = "X"
	paperMove     = "Y"
	scissorsMove  = "Z"
)

func Score(line string) int {
	moves := strings.Split(line, " ")
	game := gameState(moves[0])
	score := make(chan int)
	go game.Fire(moves[1], score)

	return <-score
}

func gameState(shape string) *stateless.StateMachine {
	gameState := stateless.NewStateMachine(shape)
	gameState.Configure(rockState).
		Permit(rockMove, drawState).
		Permit(paperMove, winState).
		Permit(scissorsMove, loseState)
	gameState.Configure(paperState).
		Permit(rockMove, loseState).
		Permit(paperMove, drawState).
		Permit(scissorsMove, winState)
	gameState.Configure(scissorsState).
		Permit(rockMove, winState).
		Permit(paperMove, loseState).
		Permit(scissorsMove, drawState)
	gameState.Configure(winState).
		OnEntryFrom(rockMove, func(_ context.Context, args ...interface{}) error {
			addPoints(7, args[0].(chan int))
			return nil
		}).
		OnEntryFrom(paperMove, func(_ context.Context, args ...interface{}) error {
			addPoints(8, args[0].(chan int))
			return nil
		}).
		OnEntryFrom(scissorsMove, func(_ context.Context, args ...interface{}) error {
			addPoints(9, args[0].(chan int))
			return nil
		})
	gameState.Configure(loseState).
		OnEntryFrom(rockMove, func(_ context.Context, args ...interface{}) error {
			addPoints(1, args[0].(chan int))
			return nil
		}).
		OnEntryFrom(paperMove, func(_ context.Context, args ...interface{}) error {
			addPoints(2, args[0].(chan int))
			return nil
		}).
		OnEntryFrom(scissorsMove, func(_ context.Context, args ...interface{}) error {
			addPoints(3, args[0].(chan int))
			return nil
		})
	gameState.Configure(drawState).
		OnEntryFrom(rockMove, func(_ context.Context, args ...interface{}) error {
			addPoints(4, args[0].(chan int))
			return nil
		}).
		OnEntryFrom(paperMove, func(_ context.Context, args ...interface{}) error {
			addPoints(5, args[0].(chan int))
			return nil
		}).
		OnEntryFrom(scissorsMove, func(_ context.Context, args ...interface{}) error {
			addPoints(6, args[0].(chan int))
			return nil
		})
	chanType := reflect.ChanOf(reflect.BothDir, reflect.TypeOf(0))
	gameState.SetTriggerParameters(rockMove, chanType)
	gameState.SetTriggerParameters(paperMove, chanType)
	gameState.SetTriggerParameters(scissorsMove, chanType)

	return gameState
}

func addPoints(points int, score chan int) {
	score <- points
}
