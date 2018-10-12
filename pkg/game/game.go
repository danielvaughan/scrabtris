package game

import (
	"github.com/danielvaughan/scrabtris/pkg/bag"
	"github.com/danielvaughan/scrabtris/pkg/board"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"log"
)

//Game manages the game state
type Game struct {
	logger   *log.Logger
	clock    *Clock
	bag      *bag.Bag
	board    *board.Board
	nextTile tile.Tile
	rate     int
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

func NewGame(logger *log.Logger, bag *bag.Bag, board *board.Board, view *View, tileLanded chan tile.Tile, topReached chan tile.Tile, r int) *Game {
	view.Board = board
	g := Game{
		logger: logger,
		clock: NewClock(func() {
			board.ProgressTile()
			view.refreshScreen()
		}),
		bag:      bag,
		board:    board,
		nextTile: tile.EmptyTile,
		rate:     r,
	}
	view.Game = &g
	go func(tileLanded chan tile.Tile) {
		for {
			select {
			case <-tileLanded:
				g.pickTile()
			case <-topReached:
				g.clock.gameover = true
			}
		}
	}(tileLanded)
	return &g
}
