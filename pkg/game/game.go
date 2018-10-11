package game

import (
	"github.com/danielvaughan/scrabtris/pkg/bag"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"log"
)

//Game manages the game state
type Game struct {
	logger   *log.Logger
	clock    *Clock
	bag      *bag.Bag
	board    *Board
	nextTile *tile.Tile
	rate     int
}

//Start starts the game
func (g *Game) Start() {
	g.clock.start()
	g.pickTile()
}

func (g *Game) pickTile() {
	if g.nextTile == &tile.EmptyTile {
		g.nextTile = g.bag.PickTile()
	}
	g.board.AddTile(g.nextTile)
	g.nextTile = g.bag.PickTile()
}

func (g *Game) onTileLanded(t tile.Tile) {
	//work out words
	g.pickTile()
}

func NewGame(logger *log.Logger, bag *bag.Bag, board *Board, view *View, r int) *Game {
	view.Board = board
	g := Game{
		logger: logger,
		clock: NewClock(func() {
			board.ProgressTile()
			view.refreshScreen()
		}),
		bag:      bag,
		board:    board,
		nextTile: &tile.EmptyTile,
		rate:     r,
	}
	board.Game = &g
	view.Game = &g
	return &g
}
