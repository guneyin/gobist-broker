package importer_test

import (
	"fmt"
	"github.com/guneyin/gobist-importer/pkg/broker"
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

	fmt.Println(b.String())

	file, err := os.ReadFile("testdata/garanti.csv")
	assertError(t, err)
	assertNotNil(t, file)

	ts, err := importer.Import(broker.Garanti, file)
	assertError(t, err)
	assertNotNil(t, ts)

	for _, item := range ts.Items {
		fmt.Printf("%-10s %-35s %-5d %-10.2f %-15s\n", item.Symbol, item.Date, item.Quantity, item.Price, item.Type.String())
	}
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
