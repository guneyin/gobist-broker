
# gobist-importer - GO library to parse and import supported stock brokers' transaction file

This project aims to parse transaction files for supported stock brokers. It's useful if you have several accounts and want to collect all transactions in one place.

### Supported Brokers
- Garanti BBVA Yatırım


## Installation

    $ go get github.com/guneyin/gobist-importer

## Usage and Example

### Create Importer
```go
ts, err := importer.New(vendors.Garanti, "garanti.csv")
if err != nil {
    log.Fatal(err)
}
```

### Example
```go
func main() {
    ts, err := importer.New(vendors.Garanti, "garanti.csv")
    if err != nil {
        log.Fatal(err)
    }

    for _, item := range ts.Items {
        fmt.Printf("%-10s %-35s %-5d %-10.2f %-15s\n", item.Symbol, item.Date, item.Quantity, item.Price, item.Type.String())
    }
}
``` 