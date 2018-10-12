package board

import (
	"fmt"
	"github.com/danielvaughan/scrabtris/pkg/tile"
)

const (
	width  = 10
	height = 18
)

//Board represents the current state of tiles on the board and the position of the active tile
type Board struct {
	squares    [width][height]tile.Tile
	tileRow    int
	tileCol    int
	tileLanded chan tile.Tile
	topReached chan tile.Tile
}

//State returns the current state of the board as text.
func (b *Board) State() string {
	text := ""
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			text = fmt.Sprintf("%s%s", text, string(b.squares[i][j].Letter))
		}
		text = fmt.Sprintf("%s\n", text)
	}
	return text
}

//AddTile adds a tile to the top row of the board in the middle column
func (b *Board) AddTile(t tile.Tile) {
	b.squares[b.tileCol][b.tileRow] = t
}

func (b *Board) moveTileDown(t tile.Tile) {
	b.squares[b.tileCol][b.tileRow] = tile.EmptyTile
	b.tileRow++
	b.squares[b.tileCol][b.tileRow] = t
}

func (b *Board) landTile(t tile.Tile) {
	b.tileLanded <- t
	b.tileRow = 0
	b.tileCol = width / 2
}

//ProgressTile progresses the in play tile down one row or lands the title if it can go no further
func (b *Board) ProgressTile() {
	t := b.squares[b.tileCol][b.tileRow]
	nextRowIsLast := b.tileRow == height-2
	nextRowIsTop := b.tileRow == 0
	nextSquareHasTile := b.squares[b.tileCol][b.tileRow+1].Letter != tile.EmptyTile.Letter
	if nextSquareHasTile && nextRowIsTop {
		b.moveTileDown(t)
		b.topReached <- t
		return
	}
	if nextRowIsLast && !nextSquareHasTile {
		b.moveTileDown(t)
		b.landTile(t)
		return
	}
	if nextSquareHasTile && !nextRowIsTop {
		b.landTile(t)
		return
	}
	b.moveTileDown(t)
}

//MoveTileLeft moves the current tile to the left
func (b *Board) MoveTileLeft() {
	if b.tileCol > 0 {
		nextSquareHasTile := b.squares[b.tileCol-1][b.tileRow].Letter != tile.EmptyTile.Letter
		if !nextSquareHasTile {
			t := b.squares[b.tileCol][b.tileRow]
			b.squares[b.tileCol][b.tileRow] = tile.EmptyTile
			b.tileCol--
			b.squares[b.tileCol][b.tileRow] = t
		}
	}
}

//MoveTileRight moves the current tile to the right
func (b *Board) MoveTileRight() {
	if b.tileCol < width-1 {
		nextSquareHasTile := b.squares[b.tileCol+1][b.tileRow].Letter != tile.EmptyTile.Letter
		if !nextSquareHasTile {
			t := b.squares[b.tileCol][b.tileRow]
			b.squares[b.tileCol][b.tileRow] = tile.EmptyTile
			b.tileCol++
			b.squares[b.tileCol][b.tileRow] = t
		}
	}
}

func (b *Board) Row() []tile.Tile {
	return nil // b.squares[height]
}

//NewBoard creates a board full of empty tiles
func NewBoard(tileLanded chan tile.Tile) *Board {
	board := &Board{
		tileRow:    0,
		tileCol:    width / 2,
		tileLanded: tileLanded,
	}
	var squares [width][height]tile.Tile
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			squares[i][j] = tile.EmptyTile
		}
	}
	board.squares = squares
	board.tileLanded = tileLanded
	return board
}
