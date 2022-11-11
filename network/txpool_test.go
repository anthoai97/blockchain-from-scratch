package network

import (
	"testing"

	"github.com/anthoai97/blockchain-from-scratch/core"
	"github.com/stretchr/testify/assert"
)

func TestTxPoll(t *testing.T) {
	txPool := NewTxPool()
	assert.Equal(t, txPool.Len(), 0)
}

func TestTxPoolAddTx(t *testing.T) {
	txPool := NewTxPool()
	tx := core.NewTransaction([]byte("foo"))
	assert.Nil(t, txPool.Add(tx))
	assert.Equal(t, txPool.Len(), 1)

	txx := core.NewTransaction([]byte("foo"))
	assert.Nil(t, txPool.Add(txx))
	assert.Equal(t, txPool.Len(), 1)

	txPool.Flush()
	assert.Equal(t, txPool.Len(), 0)
}
