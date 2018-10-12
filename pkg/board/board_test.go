package board_test

import (
	"fmt"
	"github.com/danielvaughan/scrabtris/pkg/board"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddTile(t *testing.T) {
	tileLanded := make(chan tile.Tile)
	board := board.NewBoard(tileLanded)
	go func(tileLanded chan tile.Tile) {
		for {
			select {
			case <-tileLanded:
				fmt.Println("Tile landed")
			}
		}
	}(tileLanded)
	beforeState := board.State()
	board.AddTile(tile.Tile{Letter: 'A', Score: 1})
	afterAState := board.State()
	assert.NotEqual(t, beforeState, afterAState)
	for i := 0; i < 17; i++ {
		board.ProgressTile()
	}
	endAState := board.State()
	assert.NotEqual(t, endAState, afterAState)
	board.AddTile(tile.Tile{Letter: 'B', Score: 1})
	afterBState := board.State()
	assert.NotEqual(t, afterAState, afterBState)
	for j := 0; j < 17; j++ {
		board.ProgressTile()
	}
	endBState := board.State()
	time.Sleep(1)
	assert.NotEqual(t, endBState, afterBState)
}
