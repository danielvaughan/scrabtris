package main

import (
	"github.com/danielvaughan/scrabtris/pkg/game"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	tiles := []game.Tile{{'A', 1}, {'B', 1}, {'C', 1}}
	tb := game.NewTileBag(tiles)
	ps := game.NewPubSub()
	b := game.NewBoard(ps)
	g := game.NewGame(tb, b, ps, 1)
	g.Start()
	game.WaitKeyInput()
}
