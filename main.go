package main

import (
	"github.com/danielvaughan/scrabtris/pkg/bag"
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

	bag := bag.NewUKBag()
	ps := game.NewPubSub()
	board := game.NewBoard(ps)
	g := game.NewGame(bag, board, ps, 1)
	g.Start()
	game.WaitKeyInput()
}
