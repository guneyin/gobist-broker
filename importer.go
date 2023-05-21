package importer

import (
	"github.com/guneyin/gobist-importer/pkg"
	"github.com/guneyin/gobist-importer/pkg/broker"
	"github.com/guneyin/gobist-importer/pkg/entity"
)

type Importer struct {
	broker pkg.BrokerAdapter
	file   string
}

func GetBrokers() []broker.Broker {
	return []broker.Broker{broker.Garanti}
}

func New(v broker.Broker, f string) (*Importer, error) {
	va, err := pkg.NewBrokerAdapter(v)
	if err != nil {
		return nil, err
	}

	return &Importer{
		broker: *va,
		file:   f,
	}, nil
}

func (i Importer) Import() (*entity.Transactions, error) {
	return i.broker.Parse(i.file)
}
