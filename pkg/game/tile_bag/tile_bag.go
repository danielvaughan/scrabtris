package tile_bag

// pickTile returns a random title from the TileBag
import (
	"math/rand"
	"time"
)

type TileBag struct {
	tiles []Tile
	rng   *rand.Rand
}

func (tg *TileBag) PickTile() Tile {
	pos := tg.rng.Intn(len(tg.tiles))
	return tg.tiles[pos]
}

func NewTileBag(tiles []Tile) *TileBag {
	seed := time.Now().UTC().UnixNano()
	t := TileBag{
		tiles: tiles,
		rng:   rand.New(rand.NewSource(seed)),
	}
	return &t
}
