// internal/model/game/game.go
package game

import "fmt"

type Game struct {
	size       int
	currentTurn int
	moves      []Move
	isFinished bool
}

func NewGame(size int) *Game {
	return &Game{
		size:       size,
		currentTurn: 0,
		moves:      make([]Move, 0),
		isFinished: false,
	}
}

func (g *Game) GetSize() int {
	return g.size
}

func (g *Game) SetSize(size int) {
	if size >= 4 && size <= 20 {
		g.size = size
	}
}

func (g *Game) GetCurrentTurn() int {
	return g.currentTurn
}

func (g *Game) AddMove(move Move) {
	g.moves = append(g.moves, move)
	g.currentTurn++
}

func (g *Game) GetMoves() []Move {
	return g.moves
}

func (g *Game) IsFinished() bool {
	return g.isFinished
}

func (g *Game) SetFinished(finished bool) {
	g.isFinished = finished
}

func (g *Game) String() string {
	return fmt.Sprintf("Game(size=%d, moves=%d, finished=%v)", g.size, len(g.moves), g.isFinished)
}