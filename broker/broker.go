package broker

type Broker struct {
	Name  string
	Title string
	Url   string
	Logo  string
}

type TBroker string

const (
	Garanti TBroker = "garanti"
	NCM     TBroker = "ncm"
)

func (b TBroker) String() string {
	return string(b)
}
