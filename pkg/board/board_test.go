package board

import (
	"fmt"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	tileLanded       = make(chan tile.Tile)
	topReached       = make(chan tile.Tile)
	tilePicked       = make(chan tile.Tile)
	tileMoved        = make(chan rune)
	refreshRequested = make(chan string)
	clockTicked      = make(chan int)
)

func TestAddTile(t *testing.T) {
	board := NewBoard(tileLanded, topReached, tilePicked, tileMoved, refreshRequested, clockTicked)
	go func(tileLanded chan tile.Tile) {
		for {
			select {
			case <-tileLanded:
				fmt.Println("Tile landed")
			case <-tilePicked:
				fmt.Println("Tile picked")
			case <-refreshRequested:
				fmt.Println("Refresh requested")
			}
		}
	}(tileLanded)
	beforeState := board.state()
	board.onTilePicked(tile.Tile{Letter: 'A', Score: 1})
	afterAState := board.state()
	assert.NotEqual(t, beforeState, afterAState)
	for i := 0; i < 17; i++ {
		board.onClockTicked()
	}
	endAState := board.state()
	assert.NotEqual(t, endAState, afterAState)
	board.onTilePicked(tile.Tile{Letter: 'B', Score: 1})
	afterBState := board.state()
	assert.NotEqual(t, afterAState, afterBState)
	for j := 0; j < 17; j++ {
		board.onClockTicked()
	}
	endBState := board.state()
	time.Sleep(1)
	assert.NotEqual(t, endBState, afterBState)
}
