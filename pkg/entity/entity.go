package entity

import "time"

type TrasnactionType int

const (
	Buy TrasnactionType = iota
	Sell
)

type Transaction struct {
	Symbol   string
	Date     time.Time
	Quantity int
	Price    float64
	Type     TrasnactionType
}

type Transactions struct {
	Items []Transaction
}

func (t Transactions) Count() int {
	return len(t.Items)
}

func (t *Transactions) Add(item Transaction) {
	t.Items = append(t.Items, item)
}

func (tt TrasnactionType) String() string {
	return [...]string{"Buy", "Sell", "Divided"}[tt]
}
