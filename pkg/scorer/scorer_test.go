package scorer_test

import (
	"github.com/danielvaughan/scrabtris/pkg/scorer"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestScoreWord(t *testing.T) {
	logger := log.New(os.Stdout, "test ", log.LstdFlags|log.Lshortfile)
	s := scorer.Scorer{Logger: logger}
	assert.Equal(t, 3, s.Score([]tile.Tile{{Letter: 'C', Score: 1}, {Letter: 'A', Score: 1}, {Letter: 'T', Score: 1}}))
}
