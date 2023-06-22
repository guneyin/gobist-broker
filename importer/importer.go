package importer

import (
	"errors"
	"fmt"
	"github.com/guneyin/gobist-broker/broker"
	"github.com/guneyin/gobist-broker/broker/garanti"
	"github.com/guneyin/gobist-broker/broker/ncm"
	"github.com/guneyin/gobist-broker/entity"
)

var (
	_ IBroker = (*garanti.Garanti)(nil)
	_ IBroker = (*ncm.NCM)(nil)
)

type IBroker interface {
	Get() *broker.Broker
	Parse(content []byte) (*entity.Transactions, error)
}

type BrokerAdapter struct {
	IBroker
}

func NewBrokerAdapter(b broker.TBroker) (*BrokerAdapter, error) {
	var v IBroker

	switch b {
	case broker.Garanti:
		v = garanti.Garanti{}
	case broker.NCM:
		v = ncm.NCM{}
	default:
		return nil, fmt.Errorf("unspported broker %s", b)
	}

	return &BrokerAdapter{v}, nil
}

func GetBrokers() []IBroker {
	return []IBroker{
		garanti.Garanti{},
		ncm.NCM{},
	}
}

func GetBrokerByName(name string) (*BrokerAdapter, error) {
	if ok := broker.TBroker(name); ok == "" {
		return nil, errors.New("UNSPPORTED_BROKER")
	}

	b := broker.TBroker(name)

	return NewBrokerAdapter(b)
}
