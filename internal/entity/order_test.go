package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testIfGetAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}

func testIfGetAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{ID: "321"}
	assert.Error(t, order.Validate(), "invalid price")
}

func testIfGetAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{ID: "321", Price: 32.0}
	assert.Error(t, order.Validate(), "invalid tax")
}

func TestWithAllValidParams(t *testing.T) {
	order := Order{ID: "321", Price: 10.0, Tax: 1.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
