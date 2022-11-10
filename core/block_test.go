package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/anthoai97/blockchain-from-scratch/crypto"
	"github.com/anthoai97/blockchain-from-scratch/types"
	"github.com/stretchr/testify/assert"
)

func TestSignBlock(t *testing.T) {
	b := randomBlock(t, 0, types.Hash{})
	privKey := crypto.GeneratePrivateKey()
	fmt.Println(b.Hash(BlockHasher{}))
	assert.Nil(t, b.Sign(privKey))

	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	b := randomBlock(t, 0, types.Hash{})
	privKey := crypto.GeneratePrivateKey()
	fmt.Println(b.Hash(BlockHasher{}))
	assert.Nil(t, b.Sign(privKey))

	assert.NotNil(t, b.Signature)
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}

func randomBlock(t *testing.T, height uint32, previousBlockHash types.Hash) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: previousBlockHash,
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}
	tx := randomTxWithSignature(t)

	return NewBlock(header, []Transaction{*tx})
}

func randomBlockWithSignature(t *testing.T, height uint32, previousBlockHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(t, height, previousBlockHash)
	tx := randomTxWithSignature(t)
	b.AddTransaction(tx)
	assert.Nil(t, b.Sign(privKey))

	return b
}
