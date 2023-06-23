package entity

type EnumBroker string

const (
	Garanti EnumBroker = "garanti"
	NCM     EnumBroker = "ncm"
)

func (t EnumBroker) String() string {
	return string(t)
}

type Info struct {
	Name      string
	Title     string
	TitleLong string
	Url       string
	Logo      string
}
