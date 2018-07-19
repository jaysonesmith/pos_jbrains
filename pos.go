package pos

import (
	"fmt"
	"strings"
)

var message string
var itemCollection = map[string]string{
	"000000000000": "$0.00",
	"1234567890":   "$1.00",
}

func onBarcode(barcode string) {
	processor(strings.TrimSpace(barcode))
}

func processor(barcode string) {
	pricer(barcode)
	outputMessage(message)
}

func pricer(barcode string) {
	if price, ok := itemCollection[barcode]; ok {
		message = price
		return
	}

	message = "Unable to locate price for barcode"
}

func outputMessage(message string) {
	// output to device/screen
	// or just print it since
	fmt.Println(message)
}

func lastMessageSent() string {
	return message
}
