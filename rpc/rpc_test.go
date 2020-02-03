package rpc

import (
	"encoding/json"
	"testing"
)

type BlockTemplateReply struct {
	Difficulty     int64  `json:"difficulty"`
	Height         int64  `json:"height"`
	Blob           string `json:"blocktemplate_blob"`
	ReservedOffset int    `json:"reserved_offset"`
	PrevHash       string `json:"prev_hash"`
}

func TestNewRPCCient(t *testing.T) {
	r, err := NewRPCCient()
	if err != nil {
		t.Fail()
	}

	_, err = r.GetBlockCount()
	if err != nil {
		t.Fail()
	}
}

func TestGetBlockTemplate(t *testing.T) {
	r, _  := NewRPCCient()

	_, err := r.RawRequest("getblocktemplate", []json.RawMessage{})
	if err != nil {
		t.Fail()
	}
}
