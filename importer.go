package importer

import (
	"errors"
	"github.com/guneyin/gobist-importer/pkg"
	"github.com/guneyin/gobist-importer/pkg/broker"
	"github.com/guneyin/gobist-importer/pkg/entity"
)

func GetBrokers() []broker.Broker {
	return []broker.Broker{broker.Garanti}
}

func GetBrokerByName(name string) (*broker.Broker, error) {
	if ok := broker.Broker(name); ok == "" {
		return nil, errors.New("UNSPPORTED_BROKER")
	}

	b := broker.Broker(name)

	return &b, nil
}

func Import(b broker.Broker, content []byte) (*entity.Transactions, error) {
	ba, err := pkg.NewBrokerAdapter(b)
	if err != nil {
		return nil, err
	}
	return ba.Parse(content)
}
