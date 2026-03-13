package card

import (
	"lesson-5/pkg/bank/types"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card

}

func Withdraw(card types.Card, amount types.Money) types.Card {
	if card.Active == false {
		return card

	} else if card.Balance < amount {
		return card
	} else if amount <= 0 {
		return card
	} else if card.Currency == types.TJS && amount > 20_000 {
		return card
	} else if card.Currency == types.USD && amount > 2000 {
		return card
	} else if card.Currency == types.RUB && amount > 165_000 {
		return card
	} else {
		card.Balance -= amount

	}
	return card


}
