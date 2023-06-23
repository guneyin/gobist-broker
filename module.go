package broker

import (
	"errors"
	"github.com/guneyin/gobist-broker/brokers/garanti"
	"github.com/guneyin/gobist-broker/brokers/ncm"
	"github.com/guneyin/gobist-broker/entity"
	"sync"
)

func init() {
	once := &sync.Once{}

	once.Do(func() {
		brokers = Brokers{
			entity.Garanti: garanti.New(),
			entity.NCM:     ncm.New(),
		}
	})
}

func GetBrokers() Brokers {
	return brokers
}

func GetBroker(b entity.EnumBroker) Broker {
	return brokers[b]
}

func GetBrokerByName(name string) (Broker, error) {
	if ok := entity.EnumBroker(name); ok == "" {
		return nil, errors.New("UNSPPORTED_BROKER")
	}

	b := entity.EnumBroker(name)

	return GetBroker(b), nil
}
