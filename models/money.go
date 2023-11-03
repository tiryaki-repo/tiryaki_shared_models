package models

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Money struct {
	Currency Currency
	Amount   decimal.Decimal
}

func (m *Money) IsValid() error {
	if m.Currency != TRY || m.Currency != USD {
		return fmt.Errorf("invalid currency type")
	}
	if m.Amount.IsNegative() {
		return fmt.Errorf("invalid amount")
	}
	return nil
}
