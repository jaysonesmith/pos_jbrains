package pos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnBarcodeOutputsAndReturnsMatchedBarcodes(t *testing.T) {
	testCases := []struct {
		name, barcode, expectedPrice string
	}{
		{name: "Twelve Zeros with new line", barcode: "000000000000\n", expectedPrice: "$0.00"},
		{name: "Number with new line and return", barcode: "1234567890\n\r", expectedPrice: "$1.00"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			onBarcode(tc.barcode)

			assert.Equal(t, tc.expectedPrice, lastPrice())
			assert.Equal(t, "", lastErrorMessage)
		})
	}
}

func TestOnBarcodeSetsErrorsForUnmatchedBarcodes(t *testing.T) {
	testCases := []struct {
		name, barcode, expectedPrice string
	}{
		{name: "Unmatched alpha", barcode: "alpha"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			onBarcode(tc.barcode)

			errorMessage := errorMessage()

			assert.Equal(t, "Unable to locate price for barcode", errorMessage)
		})
	}
}
