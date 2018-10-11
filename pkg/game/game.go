package game

import (
	"github.com/danielvaughan/scrabtris/pkg/bag"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"log"
	"time"
)

//Game manages the game state
type Game struct {
	logger *log.Logger
	clock  *Clock
	bag    *bag.Bag
	board  *Board
	rate   int
}

//Start starts the game
func (g *Game) Start() {
	g.pickTile()
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		g.tick()
	}
}

func (g *Game) pickTile() {
	tile := g.bag.PickTile()
	g.onTilePicked(tile)
}

func (g *Game) onTilePicked(tile tile.Tile) {
	g.board.AddTile(tile)
	RewriteScreen(g.board.State())
}

func (g *Game) tick() {
	g.onTick()
}

func (g *Game) onTick() {
	g.board.ProgressTile()
	RewriteScreen(g.board.State())
}

func (g *Game) onTileLanded(t tile.Tile) {
	g.logger.Printf("Tile %s landed", string(t.Letter))
}

func NewGame(logger *log.Logger, bag *bag.Bag, board *Board, r int) *Game {
	g := Game{
		logger: logger,
		clock: NewClock(func() {

		}),
		bag:   bag,
		board: board,
		rate:  r,
	}
	board.Game = g
	return &g
}
