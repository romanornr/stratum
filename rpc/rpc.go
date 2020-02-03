package rpc

import (
	"github.com/btcsuite/btcd/rpcclient"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"
)

type RPCClient struct {
	sync.RWMutex
	sickRate         int64
	successRate      int64
	Accepts          int64
	Rejects          int64
	LastSubmissionAt int64
	FailsCount       int64
	Url              *url.URL
	login            string
	password         string
	Name             string
	sick             bool
	client           *http.Client
	info             atomic.Value
}

func NewRPCCient() (*rpcclient.Client, error){
	connCfg := &rpcclient.ConnConfig{
		Host:         "127.0.0.1:5222",
		User:         "via",
		Pass:         "via",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}


