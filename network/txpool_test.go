package network

import (
	"math/rand"
	"strconv"
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

func TestSortTransaction(t *testing.T) {
	p := NewTxPool()
	txLen := 1000

	for i := 0; i < txLen; i++ {
		tx := core.NewTransaction([]byte(strconv.FormatInt(int64(i), 10)))
		tx.SetFirstSeen(int64(i * rand.Intn(10000)))
		assert.Nil(t, p.Add(tx))
	}

	assert.Equal(t, txLen, p.Len())

	txx := p.Transactions()
	for i := 0; i < len(txx)-1; i++ {
		assert.True(t, txx[i].FirstSeen() < txx[i+1].FirstSeen())
	}

}
