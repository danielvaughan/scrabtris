package game

import (
	"github.com/danielvaughan/scrabtris/pkg/game/tile_bag"
	"time"
)

type Game struct {
	tileBag *tile_bag.TileBag
	rate    int
}

func (game *Game) Start() {
	for i := 0; i < 10; i++ {
		game.pickTiles()
	}
}

func (game *Game) pickTiles() {
	tile := game.tileBag.PickTile()
	time.Sleep(1 * time.Second)
	println(string(tile.Letter))
}

func NewGame(tb *tile_bag.TileBag, r int) *Game {
	g := Game{
		tileBag: tb,
		rate:    r,
	}
	return &g
}
