package ncm

import (
	"github.com/guneyin/gobist-importer/pkg/broker"
	"github.com/guneyin/gobist-importer/pkg/entity"
)

type NCM broker.Broker

func (g NCM) Get() *broker.Broker {
	return &broker.Broker{
		Name:  broker.NCM.String(),
		Title: "NCM Investment Menkul Değerler A.Ş.",
		Url:   "https://ncminvest.com.tr/",
		Logo:  "https://ncminvest.com.tr/Resimler/5/5-logo.webp",
	}
}

func (g NCM) Parse(content []byte) (*entity.Transactions, error) {
	return nil, nil
}
