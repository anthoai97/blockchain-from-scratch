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
	privKey := crypto.GeneratePrivateKey()
	tx := randomTxWithSignature(t)
	header := &Header{
		Version:       1,
		PrevBlockHash: previousBlockHash,
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}

	b, err := NewBlock(header, []*Transaction{tx})
	assert.Nil(t, err)
	dataHash, err := CalculatorDataHash(b.Transactions)
	assert.Nil(t, err)
	b.Header.DataHash = dataHash
	assert.Nil(t, b.Sign(privKey))

	return b
}
