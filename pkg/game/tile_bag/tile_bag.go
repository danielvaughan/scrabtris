package tile_bag

// pickTile returns a random title from the TileBag
import (
	"math/rand"
	"time"
)

type Tile struct {
	letter rune
	score  int
}

type TileBag struct {
	tiles []Tile
	rng   *rand.Rand
}

func (tg *TileBag) pickTile() Tile {
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
