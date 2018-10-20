package tile

//Tile represents a tile on the board
type Tile struct {
	Letter rune
	Score  int
}

var EmptyTile = Tile{Letter: ' ', Score: 0}
