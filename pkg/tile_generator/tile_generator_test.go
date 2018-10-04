package tile_generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateTile(t *testing.T) {
	assert.Equal(t, 'A', generateTile())
}
