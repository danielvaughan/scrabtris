package board

import (
	"fmt"
	"github.com/danielvaughan/scrabtris/pkg/game/tile_bag"
	"github.com/danielvaughan/scrabtris/pkg/pub_sub"
)

type Board struct {
	width    int
	height   int
	squares  [][]rune
	tileRow  int
	tileCol  int
	eventBus pub_sub.PubSub
}

func (b *Board) Render() {
	for _, cells := range b.squares {
		fmt.Print("|")
		for _, cell := range cells {
			fmt.Print(string(cell))
		}
		fmt.Println("|")
	}
	for i := 0; i < b.width+2; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func (b *Board) AddTile(tile tile_bag.Tile) {
	b.tileRow = 0
	b.tileCol = b.width / 2
	b.squares[b.tileRow][b.tileCol] = tile.Letter
	b.Render()
}

func (b *Board) ProgressTile() {
	tile := b.squares[b.tileRow][b.tileCol]
	b.squares[b.tileRow][b.tileCol] = ' '
	if b.tileRow != b.height-1 && b.squares[b.tileRow+1][b.tileCol] == ' ' {
		b.tileRow++
	} else {
		b.eventBus.Pub(pub_sub.Event{EventType: "tile landed"})
	}
	b.squares[b.tileRow][b.tileCol] = tile
	b.Render()
}

func NewBoard(width int, height int, eventBus *pub_sub.PubSub) *Board {
	fmt.Printf("Creating a %d by %d board\n", width, height)
	rows := make([][]rune, height)
	for i := range rows {
		rows[i] = make([]rune, width)
	}
	for i, row := range rows {
		for j := range row {
			rows[i][j] = ' '
		}
	}
	return &Board{
		width:   width,
		height:  height,
		squares: rows,
		tileRow: -1,
		tileCol: -1,
	}
}
