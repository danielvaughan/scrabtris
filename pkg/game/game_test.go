package game

import (
	"github.com/danielvaughan/scrabtris/pkg/game/tile_bag"
	"testing"
)

func TestGameStart(t *testing.T) {
	tiles := []tile_bag.Tile{{'A', 1}, {'B', 1}, {'C', 1}}
	tb := tile_bag.NewTileBag(tiles)
	g := NewGame(tb, 1)
	g.Start()
}
