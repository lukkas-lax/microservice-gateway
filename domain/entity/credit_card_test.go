package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNumeroCartaoCredito(t *testing.T) {

	_, err := NovoCartaoCredito("5000000000000000", "Lucas Alves Xavier", 12, 2025, 123)
	assert.Equal(t, "Número de cartão inválido", err.Error())

	_, err = NovoCartaoCredito("5318303538835699", "Lucas Alves Xavier", 12, 2025, 123)
	assert.Nil(t, err)
}

func TestMesExpiracao(t *testing.T) {

	_, err := NovoCartaoCredito("5318303538835699", "Lucas Alves Xavier", 13, 2025, 123)
	assert.Equal(t, "Mês de expiração inválido", err.Error())

	_, err = NovoCartaoCredito("5318303538835699", "Lucas Alves Xavier", 0, 2025, 123)
	assert.Equal(t, "Mês de expiração inválido", err.Error())

	_, err = NovoCartaoCredito("5318303538835699", "Lucas Alves Xavier", 10, 2025, 123)
	assert.Nil(t, err)
}

func TestAnoExpiracao(t *testing.T) {

	anoAnterior := time.Now().AddDate(-1, 0, 0)
	_, err := NovoCartaoCredito("5318303538835699", "Lucas Alves Xavier", 9, anoAnterior.Year(), 123)
	assert.Equal(t, "Ano de expiração inválido", err.Error())

}
