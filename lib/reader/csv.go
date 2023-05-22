package reader

import (
	"encoding/csv"
	"io"
	"log"
	"strings"
)

func ReadCSV(content []byte, ignoreTitle bool) ([][]string, error) {
	s := strings.NewReader(string(content))

	c := csv.NewReader(s)
	c.Comma = ','
	c.LazyQuotes = true

	it := ignoreTitle
	var res [][]string

	for {
		line, err := c.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if it {
			it = false

			continue
		}

		res = append(res, line)
	}

	return res, nil
}
