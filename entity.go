package broker

import (
	"encoding/json"
	"github.com/guneyin/gobist-broker/brokers/garanti"
	"github.com/guneyin/gobist-broker/brokers/ncm"
	"github.com/guneyin/gobist-broker/entity"
)

var (
	_ Broker = (*garanti.Garanti)(nil)
	_ Broker = (*ncm.NCM)(nil)
)

type Broker interface {
	Info() entity.BrokerInfo
	Parse(content []byte) (*entity.Transactions, error)
}

type Brokers map[entity.EnumBroker]Broker

var brokers Brokers

func (b *Brokers) ToJSON() string {
	var bl []entity.BrokerInfo

	for _, item := range brokers {
		bl = append(bl, item.Info())
	}

	d, _ := json.MarshalIndent(bl, "", " ")

	return string(d)
}
