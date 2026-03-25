package card

import (
	"go-lessons/pkg/bank/types"
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

func Withdraw(card *types.Card, amount types.Money) {
	if card.Active == false {
		return

	} else if card.Balance < amount {
		return
	} else if amount <= 0 {
		return
	} else if card.Currency == types.TJS && amount > 20_000 {
		return
	} else if card.Currency == types.USD && amount > 2000 {
		return
	} else if card.Currency == types.RUB && amount > 165_000 {
		return

	}
	card.Balance -= amount

}

func Deposit(card *types.Card, amount types.Money) {
	if card.Active == false {
		return
	}
	if amount > 50_000 {
		return
	}

	card.Balance += amount
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {

	if card.Active == false {
		return
	}
	if card.Balance <= 0 {
		return
	}

	card.MinBalance = card.Balance

	bonus := card.MinBalance * types.Money(percent) * types.Money(daysInMonth) / types.Money(daysInYear) / 100

	if bonus > 5000 {
		bonus = 5000
	}

	card.Balance += bonus

}
