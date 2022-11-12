package core

import (
	"fmt"

	"github.com/anthoai97/blockchain-from-scratch/crypto"
	"github.com/anthoai97/blockchain-from-scratch/types"
)

type Transaction struct {
	Data []byte

	From      crypto.PublicKey
	Signature *crypto.Signature

	// cache version of data hash
	hash types.Hash
	// firstSeen is the timestamp of when this tx is first seen locally
	fristSeen int64
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		Data: data,
	}
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.From = privKey.PublicKey()
	tx.Signature = sig
	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("tx hash no signature")
	}

	if !tx.Signature.Verify(tx.From, tx.Data) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}

func (tx *Transaction) Hash(hasher Hasher[*Transaction]) types.Hash {
	if tx.hash.IsZero() {
		tx.hash = hasher.Hash(tx)
	}
	return tx.hash
}

func (tx *Transaction) SetFirstSeen(t int64) {
	tx.fristSeen = t
}

func (tx *Transaction) FirstSeen() int64 {
	return tx.fristSeen
}
