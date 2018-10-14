package game

import (
	"fmt"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"github.com/nsf/termbox-go"
	"log"
)

//Game manages the game state
type Game struct {
	logger           *log.Logger
	clock            *Clock
	tileRequested    chan bool
	tilePicked       chan tile.Tile
	nextTilePicked   chan tile.Tile
	tileMoved        chan rune
	refreshRequested chan string
	rate             int
}

//Start starts the game
func (g *Game) Start() {
	g.clock.start()
	g.tileRequested <- true
	g.waitKeyInput()
}

func (g *Game) gameOver() {
	g.clock.over()
}

func (g *Game) waitKeyInput() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 'q' || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD {
				fmt.Println("quit")
				return
			} else {
				if g.clock.lock {
					continue
				} else if g.clock.gameover {
					if ev.Key == termbox.KeySpace {
						g.Start()
					}
					continue
				} else if ev.Key == termbox.KeyArrowLeft {
					g.tileMoved <- 'l'
				} else if ev.Key == termbox.KeyArrowRight {
					g.tileMoved <- 'r'
				}
			}
		}
	}
}

func (g *Game) checkBoard() {
	/*for _, tt := range g.board.Rows(){
		g.dictionary.FindWords(tt)
	}*/
}

func NewGame(logger *log.Logger,
	tileRequested chan bool,
	tileLanded chan tile.Tile,
	topReached chan tile.Tile,
	tilePicked chan tile.Tile,
	nextTilePicked chan tile.Tile,
	tileMoved chan rune,
	refreshRequested chan string,
	clockTicked chan int,
	r int) *Game {
	g := Game{
		logger: logger,
		clock: NewClock(func() {
			clockTicked <- 0
		}),
		tileRequested:    tileRequested,
		tilePicked:       tilePicked,
		nextTilePicked:   nextTilePicked,
		tileMoved:        tileMoved,
		refreshRequested: refreshRequested,
		rate:             r,
	}
	g.handleEvents(topReached, tileLanded)
	return &g
}

func (g *Game) handleEvents(topReached chan tile.Tile, tileLanded chan tile.Tile) {
	go func() {
		for {
			select {
			case <-tileLanded:
				g.checkBoard()
				g.tileRequested <- true
			case <-topReached:
				g.gameOver()
			}
		}
	}()
}
