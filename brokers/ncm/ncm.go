package ncm

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/guneyin/gobist-broker/entity"
)

type transactionType string

const (
	HISSE_ALIS  transactionType = "A"
	HISSE_SATIS                 = "S"
)

var (
	errUnauthorized = errors.New("unauthorized")
	errUnknown      = errors.New("unknown error")
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
	client := resty.New()

	//client.SetDebug(true)

	r := new(parseDTO)

	err := json.Unmarshal(content, &r)
	if err != nil {
		return nil, err
	}

	if r.OtpCode != "" {
		err = b.login(client, r.UserName, r.Password, r.OtpCode)
		if err != nil {
			return nil, err
		}
	}

	if r.Token != "" {
		client.SetCookie(&http.Cookie{
			Name:     "api-access-token",
			Value:    r.Token,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
			SameSite: 0,
		})
	}

	dBegin, err := time.Parse(time.DateOnly, r.DateBegin)
	if err != nil {
		return nil, err
	}

	dEnd, err := time.Parse(time.DateOnly, r.DateEnd)
	if err != nil {
		return nil, err
	}

	return b.getTransactions(client, dBegin, dEnd)
}

func (b NCM) login(c *resty.Client, user, pwd, otp string) error {
	res := new(loginSuccessDTO)

	_, err := c.R().
		SetBody(loginRequestDTO{UserName: user, Password: pwd, OtpCode: otp}).
		SetResult(res).
		Post("https://online.ncminvest.com.tr/webapi/login")
	if err != nil {
		return err
	}

	if !res.Success {
		return errors.New(res.Message)
	}

	fmt.Println("access_token:", res.Data.AccessToken)

	return nil
}

func (b NCM) getTransactions(c *resty.Client, dBegin, dEnd time.Time) (*entity.Transactions, error) {
	tr := new(transactionResponseDTO)

	res, err := c.R().
		SetBody(transactionRequestDTO{SymbolName: "INT_GERCEKLESEN_ISLEMLER", MENKULNO: 0, FROMISLEMTARIHI: dBegin.Format(time.DateOnly), TOISLEMTARIHI: dEnd.Format(time.DateOnly)}).
		SetResult(tr).
		Post("https://online.ncminvest.com.tr/webapi/ApiCall/INT_GERCEKLESEN_ISLEMLER")
	if err != nil {
		return nil, err
	}

	switch res.StatusCode() {
	case http.StatusOK:
	case http.StatusUnauthorized:
		return nil, errUnauthorized
	default:
		return nil, errors.Join(errUnknown, err)
	}

	if !tr.Success {
		return nil, errors.Join(errUnknown, errors.New(tr.Message))
	}

	ts := new(entity.Transactions)
	var itemType entity.TransactionType

	for _, item := range tr.Data.R1 {
		d, err := time.Parse("2006-01-02T15:04:05", item.ISLEMTARIHI)
		if err != nil {
			return nil, err
		}

		switch transactionType(item.EMIR) {
		case HISSE_ALIS:
			itemType = entity.Buy
		case HISSE_SATIS:
			itemType = entity.Sell
		default:
			continue
		}

		ts.Add(entity.Transaction{
			Symbol:   item.MENKULKODU,
			Date:     d,
			Quantity: int(item.MIKTAR),
			Price:    item.FIYAT,
			Type:     itemType,
		})
	}

	return ts, nil
}
