package pkg

import (
	"fmt"
	"github.com/guneyin/gobist-importer/pkg/broker"

	"github.com/guneyin/gobist-importer/pkg/broker/garanti"
	"github.com/guneyin/gobist-importer/pkg/entity"
)

var (
	_ IBroker = (*garanti.Garanti)(nil)
)

type IBroker interface {
	Get() broker.Broker
	Parse(f string) (*entity.Transactions, error)
}

type BrokerAdapter struct {
	broker IBroker
}

func NewBrokerAdapter(b broker.Broker) (*BrokerAdapter, error) {
	var v IBroker

	switch b {
	case broker.Garanti:
		v = garanti.Garanti{}
	default:
		return nil, fmt.Errorf("unspported broker %s", b)
	}
	return &BrokerAdapter{broker: v}, nil
}

func (va *BrokerAdapter) Parse(f string) (*entity.Transactions, error) {
	fmt.Println(fmt.Sprintf("reading %v transactions from file: %s", va.broker.Get(), f))

	return va.broker.Parse(f)
}
