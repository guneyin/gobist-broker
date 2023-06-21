package broker

type Model struct {
	Name  string
	Title string
	Url   string
	Logo  string
}

type Broker string

const (
	Garanti Broker = "garanti"
)

func (v Broker) String() string {
	return string(v)
}
