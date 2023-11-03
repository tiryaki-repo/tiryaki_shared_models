package models

var (
	TRY = Currency{Code: "TRY", Symbol: "₺"}
	USD = Currency{Code: "USD", Symbol: "$"}
)

type Currency struct {
	Code   string
	Symbol string
}
