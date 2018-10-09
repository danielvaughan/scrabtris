package scorer

import (
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"log"
)

type Scorer struct {
	logger *log.Logger
}

func (s *Scorer) Score(tt []tile.Tile) int {
	score := 0
	for _, t := range tt {
		score = score + t.Score
	}
	return score
}

func NewScorer(logger *log.Logger) *Scorer {
	return &Scorer{logger: logger}
}
