package game

import (
	"fmt"
	"github.com/danielvaughan/scrabtris/pkg/bag"
	"github.com/danielvaughan/scrabtris/pkg/board"
	"github.com/danielvaughan/scrabtris/pkg/dictionary"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"github.com/nsf/termbox-go"
	"log"
)

//Game manages the game state
type Game struct {
	logger     *log.Logger
	clock      *Clock
	bag        *bag.Bag
	dictionary *dictionary.Dictionary
	board      *board.Board
	view       *View
	nextTile   tile.Tile
	rate       int
}

//Start starts the game
func (g *Game) Start() {
	g.clock.start()
	g.pickTile()
}

func (g *Game) pickTile() {
	if g.nextTile == tile.EmptyTile {
		g.nextTile = g.bag.PickTile()
	}
	g.board.AddTile(g.nextTile)
	g.nextTile = g.bag.PickTile()
}

func (g *Game) gameOver() {
	g.clock.over()
}

func (g *Game) WaitKeyInput() {
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
					g.left()
				} else if ev.Key == termbox.KeyArrowRight {
					g.right()
				}
			}
		}
		g.view.refreshScreen()
	}
}

func (g *Game) left() {
	g.board.MoveTileLeft()
}

func (g *Game) right() {
	g.board.MoveTileRight()
}

func (g *Game) checkBoard() {
	/*for _, tt := range g.board.Rows(){
		g.dictionary.FindWords(tt)
	}*/
}

func NewGame(logger *log.Logger,
	bag *bag.Bag,
	dictionary *dictionary.Dictionary,
	board *board.Board,
	view *View,
	tileLanded chan tile.Tile,
	topReached chan tile.Tile,
	r int) *Game {
	view.Board = board
	g := Game{
		logger: logger,
		clock: NewClock(func() {
			board.ProgressTile()
			view.refreshScreen()
		}),
		bag:        bag,
		dictionary: dictionary,
		board:      board,
		view:       view,
		nextTile:   tile.EmptyTile,
		rate:       r,
	}
	view.Game = &g
	go func(tileLanded chan tile.Tile) {
		for {
			select {
			case <-tileLanded:
				g.checkBoard()
				g.pickTile()
			case <-topReached:
				g.gameOver()
			}
		}
	}(tileLanded)
	return &g
}
