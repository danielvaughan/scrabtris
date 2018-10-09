package bag

import (
	"fmt"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPickingTiles(t *testing.T) {
	tiles := []tile.Tile{{'A', 1}, {'B', 1}, {'C', 1}}
	tileCounts := []TileCount{
		{1, tile.Tile{Letter: 'A', Score: 1}},
		{1, tile.Tile{Letter: 'B', Score: 1}},
		{1, tile.Tile{Letter: 'C', Score: 1}},
	}
	b := NewBag(tileCounts)
	assert.True(t, b.TileCount() == 3)
	for i := 0; i < 10; i++ {
		pickTile(b, tiles, t)
	}
}

func TestPickingUKTiles(t *testing.T) {
	b := NewUKBag()
	assert.True(t, b.TileCount() == 100)
	for i := 0; i < 1000; i++ {
		pickedTile := b.PickTile()
		t.Log(fmt.Sprintf("Picked: %s", string(pickedTile.Letter)))
		if pickedTile.Letter != '_' {
			assert.True(t, pickedTile.Score > 0)
		}
		assert.True(t, pickedTile.Letter > 0)
	}
}

func pickTile(b *Bag, tiles []tile.Tile, t *testing.T) {
	pickedTile := b.PickTile()
	t.Log(fmt.Sprintf("Picked: %s", string(pickedTile.Letter)))
	found := false
	for _, tile := range tiles {
		if pickedTile == tile {
			found = true
		}
	}
	assert.True(t, found)
}
