package main

import (
	"github.com/danielvaughan/scrabtris/pkg/game"
	"github.com/danielvaughan/scrabtris/pkg/game/board"
	"github.com/danielvaughan/scrabtris/pkg/game/tile_bag"
	"github.com/danielvaughan/scrabtris/pkg/pub_sub"
)

func main() {
	tiles := []tile_bag.Tile{{'A', 1}, {'B', 1}, {'C', 1}}
	tb := tile_bag.NewTileBag(tiles)
	ps := pub_sub.NewPubSub()
	b := board.NewBoard(9, 10, ps)
	b.Render()
	g := game.NewGame(tb, b, ps, 1)
	g.Start()
}
