package game

import (
	"github.com/danielvaughan/scrabtris/pkg/game/board"
	"github.com/danielvaughan/scrabtris/pkg/game/tile_bag"
	"time"
)

type Game struct {
	tileBag *tile_bag.TileBag
	board   *board.Board
	rate    int
}

func (g *Game) Start() {
	for i := 0; i < 10; i++ {
		g.pickTiles()
	}
}

func (g *Game) pickTiles() {
	tile := g.tileBag.PickTile()
	time.Sleep(1 * time.Second)
	g.tilePicked(tile)
}

func (g *Game) tilePicked(tile tile_bag.Tile) {
	g.board.AddTile(tile)
}

func NewGame(tb *tile_bag.TileBag, b *board.Board, r int) *Game {
	g := Game{
		tileBag: tb,
		board:   b,
		rate:    r,
	}
	return &g
}
