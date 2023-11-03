package models

type PlaceTransactionTypeId int

const (
	// add balance to place when customer buy stuff with barcode
	// this is the most primitive type of payment
	AddBalanceWithoutPaymentV1    PlaceTransactionTypeId = PlaceTransactionTypeId(1000)
	DeductBalanceWithoutPaymentV1 PlaceTransactionTypeId = PlaceTransactionTypeId(1010)
)
