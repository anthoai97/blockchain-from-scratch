package core

import (
	"fmt"
	"testing"

	"github.com/anthoai97/blockchain-from-scratch/types"
	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(t, 0, types.Hash{}))
	assert.Nil(t, err)

	return bc
}

func getPreviousBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	previousHeader, err := bc.GetHeader(height - 1)
	assert.Nil(t, err)
	return BlockHasher{}.Hash(previousHeader)
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
	assert.False(t, bc.HasBlock(1))
	assert.False(t, bc.HasBlock(100))
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlocks := 1000
	for i := 0; i < lenBlocks; i++ {
		block := randomBlockWithSignature(t, uint32(i+1), getPreviousBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(block))
	}

	assert.Equal(t, uint32(lenBlocks), bc.Height())
	assert.Equal(t, lenBlocks+1, len(bc.headers))

	assert.Error(t, bc.AddBlock(randomBlock(t, 89, types.Hash{})))
}

func TestGetHeader(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	lenBlocks := 1000
	for i := 0; i < lenBlocks; i++ {
		b := randomBlockWithSignature(t, uint32(i+1), getPreviousBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(b))
		header, err := bc.GetHeader(uint32(i + 1))
		assert.Nil(t, err)
		assert.Equal(t, header, b.Header)
	}

}

func TestAddBlocToHigh(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.Nil(t, bc.AddBlock(randomBlockWithSignature(t, 1, getPreviousBlockHash(t, bc, uint32(1)))))
	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 3, types.Hash{})))
}
