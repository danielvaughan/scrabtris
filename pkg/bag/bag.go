package bag

import (
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"math/rand"
	"time"
)

//Bag holds tiles to randomly pick from.
type Bag struct {
	tiles      []tile.Tile
	tilePicked chan tile.Tile
	rng        *rand.Rand
}

//tileCount specifies a number of tiles to add to a bag
type tileCount struct {
	number int
	tile   tile.Tile
}

func (b *Bag) tileCount() int {
	if b == nil {
		return 0
	}
	return len(b.tiles)
}

//onTileRequested returns a random tile from the bag.
func (b *Bag) onTileRequested() {
	pos := b.rng.Intn(len(b.tiles))
	b.tilePicked <- b.tiles[pos]
}

func (b *Bag) handleEvents(tileRequested chan bool) {
	go func() {
		for {
			select {
			case <-tileRequested:
				b.onTileRequested()
			}
		}
	}()
}

//NewBag creates a bag containing a specified set of tiles.
func NewBag(tileCounts []tileCount, tileRequested chan bool, tilePicked chan tile.Tile) *Bag {
	seed := time.Now().UTC().UnixNano()
	tiles := make([]tile.Tile, 0)
	for _, tc := range tileCounts {
		for i := 0; i < tc.number; i++ {
			tiles = append(tiles, tc.tile)
		}
	}
	b := Bag{
		tiles:      tiles,
		tilePicked: tilePicked,
		rng:        rand.New(rand.NewSource(seed)),
	}
	b.handleEvents(tileRequested)
	return &b
}

//NewUKBag create a bag containing the UK distribution of tiles.
func NewUKBag(tileRequested chan bool, tilePicked chan tile.Tile) *Bag {
	tileCounts := []tileCount{
		{9, tile.Tile{Letter: 'A', Score: 1}},
		{2, tile.Tile{Letter: 'B', Score: 3}},
		{2, tile.Tile{Letter: 'C', Score: 3}},
		{4, tile.Tile{Letter: 'D', Score: 2}},
		{12, tile.Tile{Letter: 'E', Score: 1}},
		{2, tile.Tile{Letter: 'F', Score: 4}},
		{3, tile.Tile{Letter: 'G', Score: 2}},
		{2, tile.Tile{Letter: 'H', Score: 4}},
		{9, tile.Tile{Letter: 'I', Score: 1}},
		{1, tile.Tile{Letter: 'J', Score: 8}},
		{1, tile.Tile{Letter: 'K', Score: 5}},
		{4, tile.Tile{Letter: 'L', Score: 1}},
		{2, tile.Tile{Letter: 'M', Score: 3}},
		{6, tile.Tile{Letter: 'N', Score: 1}},
		{8, tile.Tile{Letter: 'O', Score: 1}},
		{2, tile.Tile{Letter: 'P', Score: 3}},
		{1, tile.Tile{Letter: 'Q', Score: 10}},
		{6, tile.Tile{Letter: 'R', Score: 1}},
		{4, tile.Tile{Letter: 'S', Score: 1}},
		{6, tile.Tile{Letter: 'T', Score: 1}},
		{4, tile.Tile{Letter: 'U', Score: 1}},
		{2, tile.Tile{Letter: 'V', Score: 4}},
		{2, tile.Tile{Letter: 'W', Score: 4}},
		{1, tile.Tile{Letter: 'X', Score: 8}},
		{2, tile.Tile{Letter: 'Y', Score: 4}},
		{1, tile.Tile{Letter: 'Z', Score: 10}},
		{2, tile.Tile{Letter: '_', Score: 0}}}
	return NewBag(tileCounts, tileRequested, tilePicked)
}
