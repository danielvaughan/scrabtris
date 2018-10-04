package tile_bag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateTile(t *testing.T) {
	tiles := []Tile{{'A', 1}, {'B', 1}, {'C', 1}}
	tg := NewTileBag(tiles)
	gt := tg.pickTile()
	found := false
	for _, tile := range tiles {
		if gt == tile {
			found = true
		}
	}
	assert.True(t, found)
}
