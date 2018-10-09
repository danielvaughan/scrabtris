package dictionary

import (
	"fmt"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

//Dictionary find valid works in an array of tiles
type Dictionary struct {
	words  []string
	logger *log.Logger
}

//FindWords, given an array of tiles, returns and array of tile arrays that are valid words
func (d *Dictionary) FindWords(tt []tile.Tile) [][]tile.Tile {
	var row string
	for _, t := range tt {
		row = fmt.Sprint(row, strings.ToUpper(string(t.Letter)))
	}
	log.Printf("checking %s", row)
	for _, w := range d.words {
		log.Printf("\tagainst %s", w)
		if w == row {
			log.Printf("matched %s", w)
			return [][]tile.Tile{tt}
		}
	}
	return [][]tile.Tile{}
}

//WordCount returns the number of words in the dictionary.
func (d *Dictionary) WordCount() int {
	return len(d.words)
}

//NewDictionary returns a dictionary with the words contained in the provided wordList
func NewDictionary(logger *log.Logger, wordList io.Reader) *Dictionary {
	bytes, err := ioutil.ReadAll(wordList)
	if err != nil {
		logger.Fatal("cannot read word list.")
	}
	words := make([]string, 0)
	var word string
	for _, b := range bytes {
		if b == '\n' {
			word = strings.ToUpper(string(word))
			words = append(words, word)
			logger.Printf("added %s", word)
			word = ""
		} else {
			word = fmt.Sprint(word, string(b))
		}
	}
	f := Dictionary{
		logger: logger,
		words:  words,
	}
	return &f
}
