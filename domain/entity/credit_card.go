package entity

import (
	"errors"
	"regexp"
	"time"
)

type CartaoCredito struct {
	numero       string
	nome         string
	mesExpiracao int
	anoExpiracao int
	cvv          int
}

func NewCartaoCredito(numero string, nome string, mesExpiracao int, anoExpiracao int, cvv int) (*CartaoCredito, error) {
	cc := &CartaoCredito{
		numero:       numero,
		nome:         nome,
		mesExpiracao: mesExpiracao,
		anoExpiracao: anoExpiracao,
		cvv:          cvv,
	}

	err := cc.IsValid()
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *CartaoCredito) IsValid() error {
	err := c.validarNumero()
	if err != nil {
		return err
	}
	err = c.validarMes()
	if err != nil {
		return err
	}
	err = c.validarAno()
	if err != nil {
		return err
	}
	return nil
}

func (c *CartaoCredito) validarNumero() error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)
	if !re.MatchString(c.numero) {
		return errors.New("Número de cartão inválido")
	}
	return nil
}

func (c *CartaoCredito) validarMes() error {
	if c.mesExpiracao > 0 && c.mesExpiracao < 13 {
		return nil
	}
	return errors.New("Mês de expiração inválido")
}

func (c *CartaoCredito) validarAno() error {
	if c.anoExpiracao >= time.Now().Year() {
		return nil
	}
	return errors.New("Ano de expiração inválido")
}
