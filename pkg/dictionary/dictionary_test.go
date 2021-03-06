package dictionary_test

import (
	"github.com/danielvaughan/scrabtris/pkg/dictionary"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

func TestFindWords(t *testing.T) {
	logger := log.New(os.Stdout, "test ", log.LstdFlags|log.Lshortfile)
	d := dictionary.NewDictionary(logger, strings.NewReader("cat\ndog\ndonkey\n"))
	assert.Equal(t, 3, d.WordCount())
	ww := d.FindWords([]tile.Tile{{Letter: 'C', Score: 1}, {Letter: 'A', Score: 1}, {Letter: 'T', Score: 1}})
	assert.Equal(t, 1, len(ww))
}
