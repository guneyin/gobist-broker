package importer

import (
	"errors"
	"github.com/guneyin/gobist-importer/pkg"
	"github.com/guneyin/gobist-importer/pkg/broker"
	"github.com/guneyin/gobist-importer/pkg/broker/garanti"
	"github.com/guneyin/gobist-importer/pkg/broker/ncm"
	"github.com/guneyin/gobist-importer/pkg/entity"
)

func GetBrokers() []pkg.IBroker {
	return []pkg.IBroker{
		garanti.Garanti{},
		ncm.NCM{},
	}
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
