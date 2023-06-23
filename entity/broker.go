package entity

import "encoding/json"

type EnumBroker string

const (
	Garanti EnumBroker = "garanti"
	NCM     EnumBroker = "ncm"
)

func (t EnumBroker) String() string {
	return string(t)
}

type BrokerInfo struct {
	Name      string `json:"name"`
	Title     string `json:"title"`
	TitleLong string `json:"title_long"`
	Url       string `json:"url"`
	Logo      string `json:"logo"`
}

func (bi BrokerInfo) ToJSON() string {
	b, _ := json.MarshalIndent(&bi, "", " ")

	return string(b)
}
