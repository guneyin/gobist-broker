package broker

type Broker string

const (
	Garanti Broker = "garanti"
)

func (v Broker) String() string {
	return string(v)
}
