package board

import (
	"fmt"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddTile(t *testing.T) {
	tileLanded := make(chan tile.Tile)
	topReached := make(chan tile.Tile)
	tilePicked := make(chan tile.Tile)
	tileMoved := make(chan rune)
	clockTicked := make(chan int)
	board := NewBoard(tileLanded, topReached, tilePicked, tileMoved, clockTicked)
	go func(tileLanded chan tile.Tile) {
		for {
			select {
			case <-tileLanded:
				fmt.Println("Tile landed")
			}
		}
	}(tileLanded)
	beforeState := board.State()
	board.onTilePicked(tile.Tile{Letter: 'A', Score: 1})
	afterAState := board.State()
	assert.NotEqual(t, beforeState, afterAState)
	for i := 0; i < 17; i++ {
		board.onClockTicked()
	}
	endAState := board.State()
	assert.NotEqual(t, endAState, afterAState)
	board.onTilePicked(tile.Tile{Letter: 'B', Score: 1})
	afterBState := board.State()
	assert.NotEqual(t, afterAState, afterBState)
	for j := 0; j < 17; j++ {
		board.onClockTicked()
	}
	endBState := board.State()
	time.Sleep(1)
	assert.NotEqual(t, endBState, afterBState)
}
