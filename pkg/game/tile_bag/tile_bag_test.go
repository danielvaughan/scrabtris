package tile_bag

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPickingTiles(t *testing.T) {
	tiles := []Tile{{'A', 1}, {'B', 1}, {'C', 1}}
	tg := NewTileBag(tiles)
	for i := 0; i < 10; i++ {
		pickTile(tg, tiles, t)
	}
}

func pickTile(tg *TileBag, tiles []Tile, t *testing.T) {
	pickedTile := tg.PickTile()
	t.Log(fmt.Sprintf("Picked: %s", string(pickedTile.Letter)))
	found := false
	for _, tile := range tiles {
		if pickedTile == tile {
			found = true
		}
	}
	assert.True(t, found)
}
