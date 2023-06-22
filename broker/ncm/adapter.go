package ncm

import (
	"github.com/guneyin/gobist-broker/broker"
	"github.com/guneyin/gobist-broker/entity"
)

type NCM struct {
	info broker.Info
}

func New() *NCM {
	return &NCM{
		info: broker.Info{
			Enum:  broker.NCM,
			Name:  "NCM Investment",
			Title: "NCM Investment Menkul Değerler A.Ş.",
			Url:   "https://ncminvest.com.tr/",
			Logo:  "https://ncminvest.com.tr/Resimler/5/5-logo.webp",
		}}
}

func (b NCM) Info() broker.Info {
	return b.info
}

func (b NCM) Parse(content []byte) (*entity.Transactions, error) {
	return nil, nil
}
