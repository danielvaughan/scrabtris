package main

import (
	"github.com/danielvaughan/scrabtris/pkg/game"
	"github.com/danielvaughan/scrabtris/pkg/game/board"
	"github.com/danielvaughan/scrabtris/pkg/game/tile_bag"
)

func main() {
	tiles := []tile_bag.Tile{{'A', 1}, {'B', 1}, {'C', 1}}
	tb := tile_bag.NewTileBag(tiles)
	b := board.NewBoard(7, 10)
	b.Render()
	g := game.NewGame(tb, b, 1)
	g.Start()
}
