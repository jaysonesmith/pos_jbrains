package pos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnBarcode(t *testing.T) {
	testCases := []struct {
		name, barcode, expectedPrice string
	}{
		{name: "Twelve Zeros", barcode: "000000000000", expectedPrice: "$0.00"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			onBarcode(tc.barcode)

			actualPrice := lastPrice()

			assert.Equal(t, tc.expectedPrice, actualPrice)
		})
	}
}
