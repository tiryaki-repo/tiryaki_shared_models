package models

type CustomerTransactionTypeId int

const (
	// add balance to place when customer use bonues to buy stuff
	CustomerAddBalanceV1 CustomerTransactionTypeId = CustomerTransactionTypeId(1000)
	// deduct balance to place when customer pay stuff with its money
	CustomerDeductBalanceV1 CustomerTransactionTypeId = CustomerTransactionTypeId(1010)
)
