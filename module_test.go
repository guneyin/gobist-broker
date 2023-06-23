package broker_test

import (
	"fmt"
	"github.com/guneyin/gobist-broker"
	"github.com/guneyin/gobist-broker/entity"
	"os"
	"reflect"
	"testing"
)

func TestImporter(t *testing.T) {
	brokers := broker.GetBrokers()
	assertNotNil(t, brokers)

	fmt.Println("Supported Brokers:")

	i := 0
	for _, v := range brokers {
		i++
		fmt.Printf("	%d- %s\n", i, v.Info().TitleLong)
	}

	fmt.Println()

	b, err := broker.GetBrokerByName("garanti")
	assertError(t, err)
	assertNotNil(t, b)

	ts, err := importFile(b, "single")
	assertError(t, err)
	assertNotNil(t, ts)

	ts, err = importFile(b, "duplicated")
	assertError(t, err)
	assertNotNil(t, ts)

	ts, err = importFile(b, "full")
	assertError(t, err)
	assertNotNil(t, ts)

	fmt.Println("Imported Transactions:")
	for _, item := range ts.Items {
		fmt.Printf("	%-10s %-35s %-5d %-10.2f %-15s\n", item.Symbol, item.Date, item.Quantity, item.Price, item.Type.String())
	}
}

func importFile(b broker.Broker, t string) (*entity.Transactions, error) {
	fPath := fmt.Sprintf("testdata/%s/%s.csv", b.Info().Name, t)

	fileContent, err := os.ReadFile(fPath)
	if err != nil {
		return nil, err
	}

	return b.Parse(fileContent)
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
