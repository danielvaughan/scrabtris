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

var (
	tileRequested    = make(chan bool)
	tilePicked       = make(chan tile.Tile)
	tileLanded       = make(chan tile.Tile)
	topReached       = make(chan tile.Tile)
	nextTilePicked   = make(chan tile.Tile)
	tileMoved        = make(chan rune)
	refreshRequested = make(chan string)
	clockTicked      = make(chan int)
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

	bag.NewUKBag(tileRequested, tilePicked)
	board.NewBoard(tileLanded, topReached, tilePicked, tileMoved, refreshRequested, clockTicked)
	view.NewView(nextTilePicked, refreshRequested)
	dictionary.NewDictionary(logger, strings.NewReader("cat\ndog\ndonkey\n"))

	g := game.NewGame(logger, tileRequested, tileLanded, topReached, tilePicked, nextTilePicked, tileMoved, refreshRequested, clockTicked, 1)
	g.Start()
}
