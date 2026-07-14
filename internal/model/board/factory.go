// internal/model/board/factory.go
package board

import (
	"chess/internal/model/game"
)

func SetupInitialBoard(boardSize int) *Board {
	board := NewBoard(boardSize)
	
	// Стандартные наборы фигур
	whiteBackRow := []string{"♖", "♘", "♗", "♔", "♕", "♗", "♘", "♖"}
	whitePawns := []string{"♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙"}
	blackBackRow := []string{"♜", "♞", "♝", "♚", "♛", "♝", "♞", "♜"}
	blackPawns := []string{"♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟"}
	
	if boardSize < 8 {
		//piecesCount = boardSize
		whiteBackRow = whiteBackRow[:boardSize]
		whitePawns = whitePawns[:boardSize]
		blackBackRow = blackBackRow[:boardSize]
		blackPawns = blackPawns[:boardSize]
	}
	if boardSize > 8 {
		expand := boardSize/8 + 1
		for i := 0; i < expand; i++ {
		    whiteBackRow = append(whiteBackRow, whiteBackRow...)
			whitePawns = append(whitePawns, whitePawns...)
			blackBackRow = append(blackBackRow, blackBackRow...)
			blackPawns = append(blackPawns, blackPawns...)
	    }
		 
		
	}
	// Расставляем фигуры
	
	for i := 0;  i < boardSize; i++ {
		// Черные фигуры (сверху) - ряд 0
		if i < len(blackBackRow) {
			board.SetPiece(0, i, game.NewPiece("rook", game.Black, []rune(blackBackRow[i])[0]))
		}
		
		// Черные пешки - ряд 1
		if i < len(blackPawns) && boardSize > 2 {
			board.SetPiece(1, i, game.NewPiece("pawn", game.Black, []rune(blackPawns[i])[0]))
		}
		
		// Белые пешки - ряд boardSize-2
		if i < len(whitePawns) && boardSize > 2 {
			board.SetPiece(boardSize-2, i, game.NewPiece("pawn", game.White, []rune(whitePawns[i])[0]))
		}
		
		// Белые фигуры (снизу) - ряд boardSize-1
		if i < len(whiteBackRow) {
			board.SetPiece(boardSize-1, i, game.NewPiece("rook", game.White, []rune(whiteBackRow[i])[0]))
		}
	}
	
	return board
}