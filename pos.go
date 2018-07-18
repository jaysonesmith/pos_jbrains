package pos

import "fmt"

var itemCollection = map[string]string{
	"000000000000": "$0.00",
}
var lastPriceOutput string

func onBarcode(barcode string) {
	pricer(barcode)
}

func pricer(barcode string) {
	price := itemCollection[barcode]
	if price != "" {
		fmt.Println(price)
		lastPriceOutput = price
	}
}

func lastPrice() string {
	return lastPriceOutput
}
