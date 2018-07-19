package pos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnBarcode(t *testing.T) {
	testCases := []struct {
		name, barcode, expectedMessage string
	}{
		{name: "Twelve Zeros with new line", barcode: "000000000000\n\r", expectedMessage: "$0.00"},
		{name: "Number with new line and return", barcode: "1234567890\n\r", expectedMessage: "$1.00"},
		{name: "Unmatched alpha", barcode: "alpha\n\r", expectedMessage: "Unable to locate price for barcode"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			onBarcode(tc.barcode)
			assert.Equal(t, tc.expectedMessage, lastMessageSent())
		})
	}
}
