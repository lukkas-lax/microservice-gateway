package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("5000000000000000", "Lucas Alves Xavier", 12, 2025, 123)
	assert.Equal(t, "Invalid credit card number", err.Error())

	_, err = NewCreditCard("5318303538835699", "Lucas Alves Xavier", 12, 2025, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {

	_, err := NewCreditCard("5318303538835699", "Lucas Alves Xavier", 13, 2025, 123)
	assert.Equal(t, "Invalid expiration month", err.Error())

	_, err = NewCreditCard("5318303538835699", "Lucas Alves Xavier", 0, 2025, 123)
	assert.Equal(t, "Invalid expiration month", err.Error())

	_, err = NewCreditCard("5318303538835699", "Lucas Alves Xavier", 10, 2025, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {

	lastYear := time.Now().AddDate(-1, 0, 0)
	_, err := NewCreditCard("5318303538835699", "Lucas Alves Xavier", 9, lastYear.Year(), 123)
	assert.Equal(t, "Invalid expiration year", err.Error())

}
