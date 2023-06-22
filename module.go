package module

import (
	"errors"
	"github.com/guneyin/gobist-broker/broker"
	"github.com/guneyin/gobist-broker/broker/garanti"
	"github.com/guneyin/gobist-broker/broker/ncm"
	"github.com/guneyin/gobist-broker/entity"
	"sync"
)

var (
	_ Broker = (*garanti.Garanti)(nil)
	_ Broker = (*ncm.NCM)(nil)
)

type Broker interface {
	Info() broker.Info
	Parse(content []byte) (*entity.Transactions, error)
}

type Brokers map[broker.EnumBroker]Broker

var brokers Brokers

func init() {
	once := &sync.Once{}

	once.Do(func() {
		brokers = Brokers{
			broker.Garanti: garanti.New(),
			broker.NCM:     ncm.New(),
		}
	})
}

func GetBrokers() Brokers {
	return brokers
}

func GetBroker(b broker.EnumBroker) Broker {
	return brokers[b]
}

func GetBrokerByName(name string) (Broker, error) {
	if ok := broker.EnumBroker(name); ok == "" {
		return nil, errors.New("UNSPPORTED_BROKER")
	}

	b := broker.EnumBroker(name)

	return GetBroker(b), nil
}
