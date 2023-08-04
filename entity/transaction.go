package entity

import "time"

type TransactionType int

const (
	Buy TransactionType = iota
	Sell
)

type Transaction struct {
	Symbol   string
	Date     time.Time
	Quantity int
	Price    float64
	Type     TransactionType
}

type Transactions struct {
	Broker BrokerInfo
	Items  []Transaction
}

func (t Transactions) Count() int {
	return len(t.Items)
}

func (t *Transactions) Add(item Transaction) {
	t.Items = append(t.Items, item)
}

func (tt TransactionType) String() string {
	return [...]string{"Buy", "Sell", "Divided"}[tt]
}
