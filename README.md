
# gobist-broker - GO library which contains some useful tools for supported stock brokers

- Transaction file import. It's useful if you have several accounts and want to collect all transactions in one place.
- Give stock order

### Supported Brokers
- Garanti BBVA Yatırım
- NCM Investment

## Installation

    $ go get github.com/guneyin/gobist-importer

## Usage and Example

### Create Importer
```go
imp, err := importer.New(broker.Garanti, "garanti.csv")
if err != nil {
    log.Fatal(err)
}
```

### Example
```go
func main() {
    imp, err := importer.New(broker.Garanti, "garanti.csv")
    if err != nil {
        log.Fatal(err)
    }

    ts, err := imp.Import()
    if err != nil {
        log.Fatal(err)
    }

    for _, item := range ts.Items {
        fmt.Printf("%-10s %-35s %-5d %-10.2f %-15s\n", item.Symbol, item.Date, item.Quantity, item.Price, item.Type.String())
    }
}
``` 