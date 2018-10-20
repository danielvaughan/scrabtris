package tile

//Tile represents a tile on the board
type Tile struct {
	Letter rune
	Score  int
}

//EmptyTile is used as the default tiles on the board
var EmptyTile = Tile{Letter: ' ', Score: 0}
