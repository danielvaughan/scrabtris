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
	squares          [width][height]tile.Tile
	tileRow          int
	tileCol          int
	tileLanded       chan tile.Tile
	topReached       chan tile.Tile
	tilePicked       chan tile.Tile
	tileMoved        chan rune
	refreshRequested chan string
	clockTicked      chan int
}

//State returns the current state of the board as text.
func (b *Board) state() string {
	text := ""
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			text = fmt.Sprintf("%s%s", text, string(b.squares[i][j].Letter))
		}
		text = fmt.Sprintf("%s\n", text)
	}
	return text
}

func (b *Board) onTilePicked(t tile.Tile) {
	b.squares[b.tileCol][b.tileRow] = t
	b.refreshRequested <- b.state()
}

func (b *Board) moveTileDown(t tile.Tile) {
	b.squares[b.tileCol][b.tileRow] = tile.EmptyTile
	b.tileRow++
	b.squares[b.tileCol][b.tileRow] = t
	b.refreshRequested <- b.state()
}

func (b *Board) landTile(t tile.Tile) {
	b.tileLanded <- t
	b.tileRow = 0
	b.tileCol = width / 2
}

//ProgressTile progresses the in play tile down one row or lands the title if it can go no further
func (b *Board) onClockTicked() {
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

func (b *Board) onTileMoved(direction rune) {
	if direction == 'r' {
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
	if direction == 'l' {
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
	b.refreshRequested <- b.state()
}

func (b *Board) row() []tile.Tile {
	return nil // b.squares[height]
}

func (b *Board) handleEvents(tilePicked chan tile.Tile, tileMoved chan rune, clockTicked chan int) {
	go func() {
		for {
			select {
			case t := <-tilePicked:
				b.onTilePicked(t)
			case d := <-tileMoved:
				b.onTileMoved(d)
			case <-clockTicked:
				b.onClockTicked()
			}
		}
	}()
}

//NewBoard creates a board full of empty tiles
func NewBoard(tileLanded chan tile.Tile,
	topReached chan tile.Tile,
	tilePicked chan tile.Tile,
	tileMoved chan rune,
	refreshRequested chan string,
	clockTicked chan int) *Board {
	b := &Board{
		tileRow:          0,
		tileCol:          width / 2,
		tileLanded:       tileLanded,
		topReached:       topReached,
		tilePicked:       tilePicked,
		tileMoved:        tileMoved,
		refreshRequested: refreshRequested,
		clockTicked:      clockTicked,
		squares:          initBoard(),
	}
	b.handleEvents(tilePicked, tileMoved, clockTicked)
	return b
}

func initBoard() [10][18]tile.Tile {
	var squares [width][height]tile.Tile
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			squares[i][j] = tile.EmptyTile
		}
	}
	return squares
}
