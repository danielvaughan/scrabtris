package game

import (
	"github.com/danielvaughan/scrabtris/pkg/game/board"
	"github.com/danielvaughan/scrabtris/pkg/game/tile_bag"
	"github.com/danielvaughan/scrabtris/pkg/pub_sub"
	"time"
)

type Game struct {
	tileBag *tile_bag.TileBag
	board   *board.Board
	pubSub  *pub_sub.PubSub
	rate    int
}

func (g *Game) Start() {
	g.pickTile()
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		g.pubSub.Pub(pub_sub.Event{EventType: "tick"})
		g.tick()
	}
}

func (g *Game) pickTile() {
	tile := g.tileBag.PickTile()
	g.onTilePicked(tile)
}

func (g *Game) onTilePicked(tile tile_bag.Tile) {
	g.board.AddTile(tile)
}

func (g *Game) tick() {
	g.onTick()
}

func (g *Game) onTick() {
	g.board.ProgressTile()
}

func NewGame(tb *tile_bag.TileBag, b *board.Board, eb *pub_sub.PubSub, r int) *Game {
	g := Game{
		tileBag: tb,
		board:   b,
		pubSub:  eb,
		rate:    r,
	}
	return &g
}
