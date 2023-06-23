package garanti

import (
	"errors"

	"github.com/guneyin/gobist-broker/entity"
	"github.com/guneyin/gobist-broker/lib"
	"github.com/guneyin/gobist-broker/lib/reader"
	"strconv"
	"strings"
	"time"
)

type transactionType string

const (
	HISSE_ALIS  transactionType = "Hisse Alış"
	HISSE_SATIS                 = "Hisse Satış"
)

type Garanti struct {
	info entity.BrokerInfo
}

func New() *Garanti {
	return &Garanti{
		info: entity.BrokerInfo{
			Name:      entity.Garanti.String(),
			Title:     "Garanti Yatırım",
			TitleLong: "Garanti Yatırım Menkul Kıymetler A.Ş.",
			Url:       "https://www.garantibbvayatirim.com.tr/",
			Logo:      "https://www.garantibbvayatirim.com.tr/_assets/img/logo.svg",
		}}
}

func (b Garanti) Info() entity.BrokerInfo {
	return b.info
}

func (b Garanti) Parse(content []byte) (*entity.Transactions, error) {
	data, err := reader.ReadCSV(content, true)
	if err != nil {
		return nil, err
	}

	res := &entity.Transactions{}
	pos := len(data) - 2
	for i, line := range data {
		if i == pos {
			break
		}

		item := entity.Transaction{}
		item.Symbol = line[0]

		item.Date, err = time.Parse("02.01.2006", line[1])
		if err != nil {
			return nil, lib.ErrFileParseError(i, "date", line[1])
		}

		item.Quantity, err = strconv.Atoi(line[2])
		if err != nil {
			return nil, lib.ErrFileParseError(i, "quantity", line[2])
		}

		val := strings.Replace(line[3], ",", ".", -1)
		item.Price, err = strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, lib.ErrFileParseError(i, "price", line[3])
		}

		switch transactionType(line[6]) {
		case HISSE_ALIS:
			item.Type = entity.Buy
		case HISSE_SATIS:
			item.Type = entity.Sell
		default:
			return nil, lib.ErrFileParseError(i, "transaction_type", line[6])
		}

		res.Add(item)
	}

	if len(res.Items) == 0 {
		return nil, errors.New("no record found")
	}

	return res, nil
}
