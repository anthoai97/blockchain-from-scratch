package network

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/anthoai97/blockchain-from-scratch/core"
)

type MessageType byte

const (
	MessageTypeTx MessageType = 0x1
	MessageTypeBock
)

type RPC struct {
	From    NetAddr
	Payload io.Reader
}

type Message struct {
	Header MessageType
	Data   []byte
}

func NewMessage(t MessageType, data []byte) *Message {
	return &Message{
		Header: t,
		Data:   data,
	}
}

func (msg *Message) Bytes() []byte {
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(msg)
	return buf.Bytes()
}

type RPCHandler interface {
	HandleRPC(rpc RPC) error
}

type DefaultRPCtHandler struct {
	p RPCProcessor
}

func NewDefaultRPCHandler(p RPCProcessor) *DefaultRPCtHandler {
	return &DefaultRPCtHandler{
		p: p,
	}
}

func (h *DefaultRPCtHandler) HandleRPC(rpc RPC) error {
	msg := Message{}
	if err := gob.NewDecoder(rpc.Payload).Decode(&msg); err != nil {
		return fmt.Errorf("failed to decode messsage from %s: %s", rpc.From, err)
	}

	switch msg.Header {
	case MessageTypeTx:
		tx := new(core.Transaction)
		if err := tx.Decode(core.NewGobTxDecoder(bytes.NewReader(msg.Data))); err != nil {
			return err
		}

		return h.p.ProcessTransaction(rpc.From, tx)
	default:
		return fmt.Errorf("invalid message header %x", msg.Header)
	}
}

type RPCProcessor interface {
	ProcessTransaction(NetAddr, *core.Transaction) error
}
