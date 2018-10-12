package game_test

import (
	"github.com/danielvaughan/scrabtris/pkg/game"
	"github.com/danielvaughan/scrabtris/pkg/tile"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTile(t *testing.T) {
	board := game.NewBoard()
	beforeState := board.State()
	board.AddTile(&tile.Tile{Letter: 'A', Score: 1})
	afterAState := board.State()
	t.Log(afterAState)
	assert.NotEqual(t, beforeState, afterAState)
	board.AddTile(&tile.Tile{Letter: 'B', Score: 1})
	afterBState := board.State()
	t.Log(afterBState)
	assert.NotEqual(t, afterAState, afterBState)
}
