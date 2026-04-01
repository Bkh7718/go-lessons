package types

// Money represents a monetary amount in the smallest currency unit (e.g., cents).
type Money int64

// Currency represents the type of currency (e.g., USD, EUR).
type Currency string

// Define constants for supported currencies.
const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	RUB Currency = "RUB"
)

// PAN represents a Primary Account Number, typically used for credit/debit cards.
type PAN string

type Card struct {
	ID         int
	PAN        PAN
	Balance    Money // Recently added our own type to represent the card's balance
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

type Payment struct {
	ID     int
	Amount Money
}

type PaymentSource struct {
	Type    string // card
	Number  string // PAN
	Balance Money
}
