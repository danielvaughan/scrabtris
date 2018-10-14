package main

import (
	"github.com/danielvaughan/scrabtris/pkg/bag"
	"github.com/danielvaughan/scrabtris/pkg/board"
	"github.com/danielvaughan/scrabtris/pkg/dictionary"
	"github.com/danielvaughan/scrabtris/pkg/game"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"github.com/danielvaughan/scrabtris/pkg/view"
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"strings"
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
	bg := bag.NewUKBag()

	tileLanded := make(chan tile.Tile)
	topReached := make(chan tile.Tile)
	tilePicked := make(chan tile.Tile)
	nextTilePicked := make(chan tile.Tile)
	tileMoved := make(chan rune)
	refreshRequested := make(chan string)
	clockTicked := make(chan int)

	board.NewBoard(tileLanded, topReached, tilePicked, tileMoved, refreshRequested, clockTicked)

	dic := dictionary.NewDictionary(logger, strings.NewReader("cat\ndog\ndonkey\n"))
	v := view.NewView(nextTilePicked, refreshRequested)
	g := game.NewGame(logger, bg, dic, v, tileLanded, topReached, tilePicked, nextTilePicked, tileMoved, refreshRequested, clockTicked, 1)
	g.Start()
}
