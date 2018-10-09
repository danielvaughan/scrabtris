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
	number int
	tile   tile.Tile
}

//PickTile returns a random tile from the bag.
func (b *Bag) PickTile() tile.Tile {
	pos := b.rng.Intn(len(b.tiles))
	return b.tiles[pos]
}

//NewBag creates a bag containing a specified array of tiles.
func NewBag(tileCounts []TileCount) *Bag {
	seed := time.Now().UTC().UnixNano()
	tiles := make([]tile.Tile, 0)
	for _, tc := range tileCounts {
		for i := 0; i < tc.number; i++ {
			tiles = append(tiles, tc.tile)
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
		{1, tile.Tile{Letter: 'A', Score: 1}},
		{1, tile.Tile{Letter: 'B', Score: 1}},
		{1, tile.Tile{Letter: 'C', Score: 1}},
		{1, tile.Tile{Letter: 'D', Score: 1}},
		{1, tile.Tile{Letter: 'E', Score: 1}},
		{1, tile.Tile{Letter: 'F', Score: 1}},
		{1, tile.Tile{Letter: 'G', Score: 1}},
		{1, tile.Tile{Letter: 'H', Score: 1}},
		{1, tile.Tile{Letter: 'I', Score: 1}},
		{1, tile.Tile{Letter: 'J', Score: 1}},
		{1, tile.Tile{Letter: 'K', Score: 1}},
		{1, tile.Tile{Letter: 'L', Score: 1}},
		{1, tile.Tile{Letter: 'M', Score: 1}},
		{1, tile.Tile{Letter: 'N', Score: 1}},
		{1, tile.Tile{Letter: 'O', Score: 1}},
		{1, tile.Tile{Letter: 'P', Score: 1}},
		{1, tile.Tile{Letter: 'Q', Score: 1}},
		{1, tile.Tile{Letter: 'R', Score: 1}},
		{1, tile.Tile{Letter: 'S', Score: 1}},
		{1, tile.Tile{Letter: 'T', Score: 1}},
		{1, tile.Tile{Letter: 'U', Score: 1}},
		{1, tile.Tile{Letter: 'V', Score: 1}},
		{1, tile.Tile{Letter: 'W', Score: 1}},
		{1, tile.Tile{Letter: 'X', Score: 1}},
		{1, tile.Tile{Letter: 'Y', Score: 1}},
		{1, tile.Tile{Letter: 'Z', Score: 1}},
		{1, tile.Tile{Letter: '_', Score: 1}}}
	return NewBag(tileCounts)
}
