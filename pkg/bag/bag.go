package bag

import (
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"math/rand"
	"time"
)

//Bag holds tiles to randomly pick from.
type Bag struct {
	tiles []tile.Tile
	rng   *rand.Rand
}

//TileCount specifies a number of tiles to add to a bag
type TileCount struct {
	Number int
	Tile   tile.Tile
}

func (b *Bag) TileCount() int {
	if b == nil {
		return 0
	}
	return len(b.tiles)
}

//PickTile returns a random tile from the bag.
func (b *Bag) PickTile() tile.Tile {
	pos := b.rng.Intn(len(b.tiles))
	return b.tiles[pos]
}

//NewBag creates a bag containing a specified set of tiles.
func NewBag(tileCounts []TileCount) *Bag {
	seed := time.Now().UTC().UnixNano()
	tiles := make([]tile.Tile, 0)
	for _, tc := range tileCounts {
		for i := 0; i < tc.Number; i++ {
			tiles = append(tiles, tc.Tile)
		}
	}
	t := Bag{
		tiles: tiles,
		rng:   rand.New(rand.NewSource(seed)),
	}
	return &t
}

//NewUKBag create a bag containing the UK distribution of tiles.
func NewUKBag() *Bag {
	tileCounts := []TileCount{
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
	return NewBag(tileCounts)
}
