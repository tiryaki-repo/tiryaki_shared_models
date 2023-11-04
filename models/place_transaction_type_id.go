package models

type PlaceTransactionTypeId int

const (
	// add balance to place when customer use bonues to buy stuff
	AddBalanceV1 PlaceTransactionTypeId = PlaceTransactionTypeId(1000)
	// deduct balance to place when customer pay stuff with its money
	DeductBalanceV1 PlaceTransactionTypeId = PlaceTransactionTypeId(1010)
)
