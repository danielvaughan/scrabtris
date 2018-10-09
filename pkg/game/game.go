package game

import (
	"github.com/danielvaughan/scrabtris/pkg/bag"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"time"
)

type Game struct {
	clock  *Clock
	bag    *bag.Bag
	board  *Board
	pubSub *PubSub
	rate   int
}

func (g *Game) Start() {
	g.pickTile()
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		g.pubSub.Pub(Event{EventType: "tick"})
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

func NewGame(bag *bag.Bag, board *Board, eb *PubSub, r int) *Game {
	g := Game{
		clock: NewClock(func() {

		}),
		bag:    bag,
		board:  board,
		pubSub: eb,
		rate:   r,
	}
	return &g
}
