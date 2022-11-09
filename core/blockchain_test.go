package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)

	return bc
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))

	fmt.Println(bc.Height())
}

func TestHashBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlocks := 1000
	for i := 0; i < lenBlocks; i++ {
		block := randomBlockWithSignature(t, uint32(i+1))
		assert.Nil(t, bc.AddBlock(block))
	}

	assert.Equal(t, uint32(lenBlocks), bc.Height())
	assert.Equal(t, lenBlocks+1, len(bc.headers))

	assert.Error(t, bc.AddBlock(randomBlock(89)))
}
