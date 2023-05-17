package reader

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ReadCSV(f string, it bool) ([][]string, error) {
	var res [][]string

	ignoreTitle := it

	csvfile, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	c := csv.NewReader(csvfile)
	c.Comma = ','
	c.LazyQuotes = true

	for {
		line, err := c.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if ignoreTitle {
			ignoreTitle = false

			continue
		}

		res = append(res, line)
	}

	return res, nil
}
