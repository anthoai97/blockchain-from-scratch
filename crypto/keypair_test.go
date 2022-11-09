package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeypair_Sign_Verify_Success(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	msg := []byte("Hello World")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	assert.True(t, sig.Verify(pubKey, msg))
}

func TestKeypair_Sign_Verify_Fail(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	msg := []byte("Hello World")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivKey := GeneratePrivateKey()
	otherPubKey := otherPrivKey.PublicKey()

	assert.False(t, sig.Verify(otherPubKey, msg))
	assert.False(t, sig.Verify(pubKey, []byte("Xxxxxx")))
}
