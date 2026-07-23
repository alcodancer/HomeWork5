// internal/model/board/board.go
package board

import "fmt"

type Board struct {
	cells [][]Cell
	size  int
}

type Cell struct {
	row     int
	col     int
	piece   interface{} // Можно заменить на конкретный тип Piece из пакета game
	isEmpty bool
}

func NewBoard(size int) *Board {
	cells := make([][]Cell, size)
	for i := 0; i < size; i++ {
		cells[i] = make([]Cell, size)
		for j := 0; j < size; j++ {
			cells[i][j] = Cell{
				row:     i,
				col:     j,
				isEmpty: true,
				piece:   nil,
			}
		}
	}
	return &Board{
		cells: cells,
		size:  size,
	}
}

func (b *Board) GetSize() int {
	return b.size
}

func (b *Board) GetCell(row, col int) *Cell {
	if row >= 0 && row < b.size && col >= 0 && col < b.size {
		return &b.cells[row][col]
	}
	return nil
}

func (b *Board) SetPiece(row, col int, piece interface{}) {
	if cell := b.GetCell(row, col); cell != nil {
		cell.piece = piece
		cell.isEmpty = (piece == nil)
	}
}

func (b *Board) IsEmpty(row, col int) bool {
	if cell := b.GetCell(row, col); cell != nil {
		return cell.isEmpty
	}
	return false
}

func (b *Board) GetPiece(row, col int) interface{} {
	if cell := b.GetCell(row, col); cell != nil {
		return cell.piece
	}
	return nil
}

func (b *Board) String() string {
	result := "  "
	for i := 0; i < b.size; i++ {
		result += fmt.Sprintf("%c|", 'A'+i)
	}
	result += "\n"
	
	for i := 0; i < b.size; i++ {
		result += fmt.Sprintf("%d ", i+1)
		for j := 0; j < b.size; j++ {
			cell := b.cells[i][j]
			if !cell.isEmpty && cell.piece != nil {
				// Пытаемся привести к типу Piece из пакета game
				switch p := cell.piece.(type) {
				case interface{ GetSymbol() rune }:
					result += fmt.Sprintf("%c|", p.GetSymbol())
				default:
					result += fmt.Sprintf("%v|", cell.piece)
				}
			} else if (i+j)%2 == 0 {
				result += " |"
			} else {
				result += "▒|"
			}
		}
		result += "\n"
	}
	return result
}