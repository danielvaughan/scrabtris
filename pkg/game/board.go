package game

import (
	"fmt"
	"github.com/danielvaughan/scrabtris/pkg/tile"
)

const (
	boardWidth  = 10
	boardHeight = 18
)

//Board represents the current state of tiles on the board and the position of the active tile
type Board struct {
	squares *[boardWidth][boardHeight]tile.Tile
	tileRow int
	tileCol int
	Game    *Game
}

//State returns the current state of the board as text.
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

//AddTile adds a tile to the top row of the board in the middle column
func (b *Board) AddTile(t *tile.Tile) {
	b.tileRow = 0
	b.tileCol = boardWidth / 2
	b.squares[b.tileCol][b.tileRow] = *t
}

//ProgressTile progresses the in play tile down one row or lands the title if it can go no further
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

//NewBoard creates a board full of empty tiles
func NewBoard() *Board {
	board := &Board{
		tileRow: -1,
		tileCol: -1,
	}
	var squares [boardWidth][boardHeight]tile.Tile
	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardHeight; j++ {
			squares[i][j] = tile.EmptyTile
		}
	}
	board.squares = &squares
	return board
}
