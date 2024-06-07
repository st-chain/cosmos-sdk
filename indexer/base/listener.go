package indexerbase

import (
	"encoding/json"
)

type PhysicalListener struct {
	StartBlock    func(uint64) error
	OnBlockHeader func(BlockHeaderData) error
	OnTx          func(TxData) error
	OnEvent       func(EventData) error
	OnKVPair      func(module string, key, value []byte, delete bool) error
	Commit        func() error
}

// TODO: combine physical and logical listeners into Listener, how would ensure setup work there?
type LogicalListener struct {
	PhysicalListener
	EnsureSetup    func(LogicalSetupData) error
	OnEntityUpdate func(module string, update EntityUpdate) error
}

type LogicalSetupData struct {
	Schema Schema
}

type BlockHeaderData struct {
	Height uint64
	Bytes  ToBytes
	JSON   ToJSON
}

type TxData struct {
	TxIndex uint64
	Bytes   ToBytes
	JSON    ToJSON
}

type EventData struct {
	TxIndex    uint64
	MsgIndex   uint64
	EventIndex uint64
	Type       string
	Data       ToJSON
}

type ToBytes = func() ([]byte, error)
type ToJSON = func() (json.RawMessage, error)
