package test

import (
	"encoding/json"
	"fmt"
	broker "github.com/guneyin/gobist-broker"
	"github.com/guneyin/gobist-broker/entity"
	"reflect"
	"testing"
	"time"
)

var (
	userName = ""
	password = ""
	otpCode  = ""
	token    = ""
)

func TestImporter(t *testing.T) {
	b := broker.GetBroker(entity.NCM)
	assertNotNil(t, b)

	dEnd := time.Now()
	dBegin := dEnd.Add(time.Hour * 24 * 30 * -12)

	ts, err := fetchTransactions(b, dBegin.Format(time.DateOnly), dEnd.Format(time.DateOnly))
	assertError(t, err)
	assertNotNil(t, ts)

	fmt.Println("Imported Transactions:")
	for _, item := range ts.Items {
		fmt.Printf("	%-10s %-35s %-5d %-10.2f %-15s\n", item.Symbol, item.Date, item.Quantity, item.Price, item.Type.String())
	}
}

func fetchTransactions(b broker.Broker, dBegin, dEnd string) (*entity.Transactions, error) {
	data := map[string]string{
		"username":  userName,
		"password":  password,
		"otpCode":   otpCode,
		"dateBegin": dBegin,
		"dateEnd":   dEnd,
		"token":     token,
	}

	content, _ := json.Marshal(data)

	return b.Parse(content)
}

func assertError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
}

func assertNotNil(t *testing.T, v interface{}) {
	if isNil(v) {
		t.Errorf("[%v] was expected to be non-nil", v)
	}
}

func isNil(v interface{}) bool {
	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)
	kind := rv.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && rv.IsNil() {
		return true
	}

	return false
}
