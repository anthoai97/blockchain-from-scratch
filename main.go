package main

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"

	"github.com/anthoai97/blockchain-from-scratch/core"
	"github.com/anthoai97/blockchain-from-scratch/crypto"
	"github.com/anthoai97/blockchain-from-scratch/network"
	"github.com/sirupsen/logrus"
)

// Server
// Transport => tcp, udp,
// Block
// Transaction
// KeyPair

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			if err := sendTranasction(trRemote, trLocal.Addr()); err != nil {
				logrus.Error(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	privKey := crypto.GeneratePrivateKey()
	opts := network.ServerOpts{
		PrivateKey: &privKey,
		ID:         "LOCAL",
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}

func sendTranasction(tr network.Transport, to network.NetAddr) error {
	privKey := crypto.GeneratePrivateKey()
	data := []byte(strconv.FormatInt(int64(rand.Intn(1000000000)), 10))
	tx := core.NewTransaction(data)
	tx.Sign(privKey)
	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return err
	}

	msg := network.NewMessage(network.MessageTypeTx, buf.Bytes())
	tr.SendMessage(to, msg.Bytes())
	return nil
}
