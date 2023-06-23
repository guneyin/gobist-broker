package ncm

import (
	"github.com/guneyin/gobist-broker/entity"
)

type NCM struct {
	info entity.BrokerInfo
}

func New() *NCM {
	return &NCM{
		info: entity.BrokerInfo{
			Name:      entity.NCM.String(),
			Title:     "NCM Investment",
			TitleLong: "NCM Investment Menkul Değerler A.Ş.",
			Url:       "https://ncminvest.com.tr/",
			Logo:      "https://ncminvest.com.tr/Resimler/5/5-logo.webp",
		}}
}

func (b NCM) Info() entity.BrokerInfo {
	return b.info
}

func (b NCM) Parse(content []byte) (*entity.Transactions, error) {
	return nil, nil
}
