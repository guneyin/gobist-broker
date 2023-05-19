package vendors

type Vendor string

const (
	Garanti Vendor = "garanti"
)

func (v Vendor) String() string {
	return string(v)
}
