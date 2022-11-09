package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/anthoai97/blockchain-from-scratch/crypto"
	"github.com/anthoai97/blockchain-from-scratch/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Height:    height,
		Timestamp: time.Now().UnixNano(),
	}
	tx := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(header, []Transaction{tx})
}

func randomBlockWithSignature(t *testing.T, height uint32) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(height)
	assert.Nil(t, b.Sign(privKey))

	return b
}

func TestSignBlock(t *testing.T) {
	b := randomBlock(0)
	privKey := crypto.GeneratePrivateKey()
	fmt.Println(b.Hash(BlockHasher{}))
	assert.Nil(t, b.Sign(privKey))

	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	b := randomBlock(0)
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
