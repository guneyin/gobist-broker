package garanti

import (
	"github.com/guneyin/gobist-importer/lib"
	"github.com/guneyin/gobist-importer/lib/reader"
	"github.com/guneyin/gobist-importer/pkg/entity"
	"github.com/guneyin/gobist-importer/pkg/vendors"
	"strconv"
	"strings"
	"time"
)

type transactionType string

const (
	HISSE_ALIS  transactionType = "Hisse Alış"
	HISSE_SATIS                 = "Hisse Satış"
)

type Garanti struct{}

func (g Garanti) Get() vendors.Vendor {
	return vendors.Garanti
}

func (g Garanti) Parse(f string) (*entity.Transactions, error) {
	data, err := reader.ReadCSV(f, true)
	if err != nil {
		return nil, err
	}

	res := &entity.Transactions{}
	pos := len(data) - 3
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

	return res, nil
}
