package broker

type EnumBroker string

const (
	Garanti EnumBroker = "garanti"
	NCM     EnumBroker = "ncm"
)

func (t EnumBroker) String() string {
	return string(t)
}

type Info struct {
	Enum  EnumBroker
	Name  string
	Title string
	Url   string
	Logo  string
}
