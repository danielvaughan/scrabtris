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
	squares  [boardWidth][boardHeight]rune
	tileRow  int
	tileCol  int
	eventBus PubSub
}

func (b *Board) State() string {
	text := ""
	for j := 0; j < boardHeight; j++ {
		for i := 0; i < boardWidth; i++ {
			text = fmt.Sprintf("%s%s", text, string(b.squares[i][j]))
		}
		text = fmt.Sprintf("%s\n", text)
	}
	return text
}

func (b *Board) AddTile(tile tile.Tile) {
	b.tileRow = 0
	b.tileCol = boardWidth / 2
	b.squares[b.tileCol][b.tileRow] = tile.Letter
}

func (b *Board) ProgressTile() {
	tile := b.squares[b.tileCol][b.tileRow]
	b.squares[b.tileCol][b.tileRow] = ' '
	if b.tileRow != boardHeight-1 && b.squares[b.tileCol][b.tileRow+1] == ' ' {
		b.tileRow++
	} else {
		b.eventBus.Pub(Event{EventType: "tile landed"})
	}
	b.squares[b.tileCol][b.tileRow] = tile
}

func NewBoard(eventBus *PubSub) *Board {
	board := &Board{
		tileRow: -1,
		tileCol: -1,
	}
	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardHeight; j++ {
			board.squares[i][j] = ' '
		}
	}
	return board
}
