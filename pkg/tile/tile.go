package tile

type Tile struct {
	Letter rune
	Score  int
}

var EmptyTile = Tile{Letter: ' ', Score: 0}
