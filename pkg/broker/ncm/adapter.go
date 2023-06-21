package ncm

import (
	"github.com/guneyin/gobist-importer/pkg/broker"
	"github.com/guneyin/gobist-importer/pkg/entity"
)

type transactionType string

const (
	HISSE_ALIS  transactionType = "Hisse Alış"
	HISSE_SATIS                 = "Hisse Satış"
)

type NCM struct {
	broker.Model
}

func (g NCM) Get() broker.Model {
	return broker.Model{
		Name:  broker.NCM.String(),
		Title: "NCM Investment Menkul Değerler A.Ş.",
		Url:   "https://ncminvest.com.tr/",
		Logo:  "https://ncminvest.com.tr/Resimler/5/5-logo.webp",
	}
}

func (g NCM) Parse(content []byte) (*entity.Transactions, error) {
	return nil, nil
}
