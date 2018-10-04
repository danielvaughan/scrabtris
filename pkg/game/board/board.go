package board

import (
	"fmt"
	"github.com/danielvaughan/scrabtris/pkg/game/tile_bag"
)

type Board struct {
	width   int
	height  int
	squares [][]rune
}

func (b *Board) Render() {
	for _, cells := range b.squares {
		fmt.Println()
		for _, cell := range cells {
			fmt.Print(string(cell))
		}
	}
	fmt.Println()
}

func (b *Board) AddTile(tile tile_bag.Tile) {
	fmt.Printf(string(b.squares[0]))
	b.Render()
}

func NewBoard(width int, height int) *Board {
	fmt.Printf("Creating a %d by %d board\n", width, height)
	rows := make([][]rune, height)
	for i := range rows {
		rows[i] = make([]rune, width)
	}
	return &Board{
		width:   width,
		height:  height,
		squares: rows,
	}
}
