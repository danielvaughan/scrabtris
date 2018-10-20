package bag

import (
	"fmt"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	tileRequested = make(chan bool)
	tilePicked    = make(chan tile.Tile)
)

func TestPickingTiles(t *testing.T) {
	tiles := []tile.Tile{{Letter: 'A', Score: 1}, {Letter: 'B', Score: 1}, {Letter: 'C', Score: 1}}
	tileCounts := []tileCount{
		{1, tile.Tile{Letter: 'A', Score: 1}},
		{1, tile.Tile{Letter: 'B', Score: 1}},
		{1, tile.Tile{Letter: 'C', Score: 1}},
	}

	b := NewBag(tileCounts, tileRequested, tilePicked)
	assert.True(t, b.tileCount() == 3)
	for i := 0; i < 10; i++ {
		pickTile(tiles, t)
	}
}

func TestPickingUKTiles(t *testing.T) {
	b := NewUKBag(tileRequested, tilePicked)
	assert.True(t, b.tileCount() == 100)
	for i := 0; i < 1000; i++ {
		tileRequested <- true
		pickedTile := <-tilePicked
		t.Log(fmt.Sprintf("Picked: %s", string(pickedTile.Letter)))
		if pickedTile.Letter != '_' {
			assert.True(t, pickedTile.Score > 0)
		}
		assert.True(t, pickedTile.Letter > 0)
	}
}

func pickTile(tiles []tile.Tile, t *testing.T) {
	tileRequested <- true
	pickedTile := <-tilePicked
	t.Log(fmt.Sprintf("Picked: %s", string(pickedTile.Letter)))
	found := false
	for _, t := range tiles {
		if pickedTile == t {
			found = true
		}
	}
	assert.True(t, found)
}
