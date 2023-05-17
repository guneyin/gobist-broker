package pkg

import (
	"fmt"
	"importer/pkg/entity"

	"importer/pkg/vendors"
	"importer/pkg/vendors/garanti"
)

var (
	_ IVendor = (*garanti.Garanti)(nil)
)

type IVendor interface {
	Get() vendors.Vendor
	Parse(f string) (*entity.Transactions, error)
}

type VendorAdapter struct {
	vendor IVendor
}

func NewVendorAdapter(vendor vendors.Vendor) (*VendorAdapter, error) {
	var v IVendor

	switch vendor {
	case vendors.Garanti:
		v = garanti.Garanti{}
	default:
		return nil, fmt.Errorf("unspported vendor %s", vendor)
	}
	return &VendorAdapter{vendor: v}, nil
}
func (va *VendorAdapter) Parse(f string) (*entity.Transactions, error) {
	fmt.Println(fmt.Sprintf("reading %v transactions from file: %s", va.vendor.Get(), f))

	return va.vendor.Parse(f)
}
