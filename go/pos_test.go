package pos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnBarcodeOutputsAndReturnsMatchedBarcodes(t *testing.T) {
	testCases := []struct {
		name, barcode, expectedPrice string
	}{
		{name: "Twelve Zeros", barcode: "000000000000", expectedPrice: "$0.00"},
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
		{name: "Alpha", barcode: "alpha"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			onBarcode(tc.barcode)

			errorMessage := errorMessage()

			assert.Equal(t, "Unable to locate price for barcode", errorMessage)
		})
	}
}
