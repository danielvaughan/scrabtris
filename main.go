package main

import (
	"github.com/danielvaughan/scrabtris/pkg/bag"
	"github.com/danielvaughan/scrabtris/pkg/game"
	"github.com/nsf/termbox-go"
	"log"
	"os"
)

func main() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	logger := log.New(os.Stdout, "scrabtris ", log.LstdFlags|log.Lshortfile)
	bag := bag.NewUKBag()
	board := game.NewBoard()
	view := &game.View{}
	g := game.NewGame(logger, bag, board, view, 1)
	g.Start()
	game.WaitKeyInput()
}
