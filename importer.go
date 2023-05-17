package importer

import (
	"importer/pkg"
	"importer/pkg/entity"
	"importer/pkg/vendors"
)

type Importer struct {
	vendor pkg.VendorAdapter
	file   string
}

func New(v vendors.Vendor, f string) (*Importer, error) {
	va, err := pkg.NewVendorAdapter(v)
	if err != nil {
		return nil, err
	}

	return &Importer{
		vendor: *va,
		file:   f,
	}, nil
}

func (i Importer) Import() (*entity.Transactions, error) {
	return i.vendor.Parse(i.file)
}
