package pos

import (
	"fmt"
	"strings"
)

var lastPriceOutput string
var lastErrorMessage string
var itemCollection = map[string]string{
	"000000000000": "$0.00",
	"1234567890":   "$1.00",
}

// no return value
func onBarcode(barcode string) {
	pricer(strings.TrimSpace(barcode))
}

func pricer(barcode string) {
	price := itemCollection[barcode]
	if price != "" {
		// printing price as stand in for output
		// to device I don't have
		fmt.Println(price)
		lastPriceOutput = price
	} else {
		error := "Unable to locate price for barcode"
		fmt.Println(error)
		lastErrorMessage = error
	}
}

func lastPrice() string {
	return lastPriceOutput
}

func errorMessage() string {
	return lastErrorMessage
}
