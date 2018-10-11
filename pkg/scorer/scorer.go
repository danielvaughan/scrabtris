package scorer

import (
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"log"
)

//Scorer provides the score for words
type Scorer struct {
	Logger *log.Logger
}

//Score returns the score for a given array of tiles
func (s *Scorer) Score(tt []tile.Tile) int {
	score := 0
	for _, t := range tt {
		score = score + t.Score
	}
	return score
}
