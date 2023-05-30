package importer_test

import (
	"fmt"
	"github.com/guneyin/gobist-importer/pkg/broker"
	"github.com/guneyin/gobist-importer/pkg/entity"
	"os"
	"reflect"
	"testing"

	"github.com/guneyin/gobist-importer"
)

func TestImporter(t *testing.T) {
	brokers := importer.GetBrokers()
	assertNotNil(t, brokers)

	for _, b := range brokers {
		fmt.Printf("Broker Name: %s\n", b.String())
	}

	b, err := importer.GetBrokerByName("garanti")
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

	for _, item := range ts.Items {
		fmt.Printf("%-10s %-35s %-5d %-10.2f %-15s\n", item.Symbol, item.Date, item.Quantity, item.Price, item.Type.String())
	}
}

func importFile(b *broker.Broker, t string) (*entity.Transactions, error) {
	fPath := fmt.Sprintf("testdata/%s/%s.csv", b.String(), t)

	file, err := os.ReadFile(fPath)
	if err != nil {
		return nil, err
	}

	return importer.Import(broker.Garanti, file)
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
