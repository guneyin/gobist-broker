package importer_test

import (
	"fmt"
	"importer"
	"importer/pkg/vendors"
	"reflect"
	"testing"
)

func TestImporter(t *testing.T) {
	imp, err := importer.New(vendors.Garanti, "testdata/garanti.csv")

	assertError(t, err)
	assertNotNil(t, imp)

	ts, err := imp.Import()
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
