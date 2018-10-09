package dictionary

import (
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

func TestFindWords(t *testing.T) {
	logger := log.New(os.Stdout, "test ", log.LstdFlags|log.Lshortfile)
	d := NewDictionary(logger, strings.NewReader("cat\ndog\ndonkey\n"))
	assert.Equal(t, 3, d.WordCount())
	ww := d.FindWords([]tile.Tile{{'C', 1}, {'A', 1}, {'T', 1}})
	assert.Equal(t, 1, len(ww))
}
