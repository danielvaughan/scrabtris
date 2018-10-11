package game

import (
	"fmt"
	"github.com/danielvaughan/scrabtris/pkg/tile"
)

const (
	boardWidth  = 10
	boardHeight = 18
)

type Board struct {
	squares [boardWidth][boardHeight]tile.Tile
	tileRow int
	tileCol int
	Game    Game
}

func (b *Board) State() string {
	text := ""
	for j := 0; j < boardHeight; j++ {
		for i := 0; i < boardWidth; i++ {
			text = fmt.Sprintf("%s%s", text, string(b.squares[i][j].Letter))
		}
		text = fmt.Sprintf("%s\n", text)
	}
	return text
}

func (b *Board) AddTile(tile tile.Tile) {
	b.tileRow = 0
	b.tileCol = boardWidth / 2
	b.squares[b.tileCol][b.tileRow] = tile
}

func (b *Board) ProgressTile() {
	t := b.squares[b.tileCol][b.tileRow]
	b.squares[b.tileCol][b.tileRow] = tile.EmptyTile
	if b.tileRow != boardHeight-1 && b.squares[b.tileCol][b.tileRow+1].Letter == tile.EmptyTile.Letter {
		b.tileRow++
	} else {
		b.Game.onTileLanded(t)
	}
	b.squares[b.tileCol][b.tileRow] = t
}

func NewBoard() *Board {
	board := &Board{
		tileRow: -1,
		tileCol: -1,
	}
	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardHeight; j++ {
			board.squares[i][j] = tile.EmptyTile
		}
	}
	return board
}
