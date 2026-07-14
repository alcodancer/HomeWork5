// internal/model/game/move.go
package game

import "fmt"

type Move struct {
	fromRow int
	fromCol int
	toRow   int
	toCol   int
	piece   Piece
}

func NewMove(fromRow, fromCol, toRow, toCol int, piece Piece) Move {
	return Move{
		fromRow: fromRow,
		fromCol: fromCol,
		toRow:   toRow,
		toCol:   toCol,
		piece:   piece,
	}
}

func (m *Move) GetFromRow() int {
	return m.fromRow
}

func (m *Move) GetFromCol() int {
	return m.fromCol
}

func (m *Move) GetToRow() int {
	return m.toRow
}

func (m *Move) GetToCol() int {
	return m.toCol
}

func (m *Move) GetPiece() Piece {
	return m.piece
}

func (m *Move) String() string {
	return fmt.Sprintf("Move: (%d,%d) -> (%d,%d) %s", m.fromRow, m.fromCol, m.toRow, m.toCol, m.piece)
}