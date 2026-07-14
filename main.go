package main

import (
	"chess/internal/model/board"
	"chess/internal/model/game"
	"chess/internal/model/player"
	"fmt"
)

func main() {
	// Создаем игроков
	player1 := askPlayerName("первого")
	player2 := askPlayerName("второго")
	
	// Запрашиваем размер доски
	boardSize := askBoardSize()
	
	// Создаем игру
	chessGame := game.NewGame(boardSize)
	
	// Создаем доску с фигурами
	gameBoard := board.SetupInitialBoard(boardSize)
	
	// Выводим информацию
	fmt.Printf("\nИгроки: %s vs %s\n", player1.GetName(), player2.GetName())
	fmt.Printf("Размер доски: %dx%d\n", chessGame.GetSize(), chessGame.GetSize())
	fmt.Println("\nДоска:")
	fmt.Print(gameBoard.String())
}

func askPlayerName(order string) *player.Player {
	var name string
	fmt.Printf("Как зовут %s игрока? ", order)
	fmt.Scan(&name)
	return player.NewPlayer(name)
}

func askBoardSize() int {
	var size int
	fmt.Print("Какое будет размер игрового поля? ")
	fmt.Scan(&size)
	return size
}
