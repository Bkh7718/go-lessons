package card

import (
	"go-lessons/pkg/bank/types"
)

// IssueCard creates and returns a new card with the specified currency, color, and name. The card is initialized with a default ID, PAN, zero balance, and active status.
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

// Withdraw deducts the specified amount from the card's balance if the card is active, has sufficient funds, and the amount does not exceed defined limits based on the currency.
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

// Deposit adds the specified amount to the card's balance if the card is active and the amount does not exceed the defined limit.
func Deposit(card *types.Card, amount types.Money) {
	if card.Active == false {
		return
	}
	if amount > 50_000 {
		return
	}

	card.Balance += amount
}

// AddBonus calculates and adds a bonus to the card's balance based on the provided percentage and time factors.
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

func Total(cards []types.Card) types.Money {

	sum := int64(0)
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			sum += int64(card.Balance)
		}

	}

	return types.Money(sum)
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var sources []types.PaymentSource
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			source := types.PaymentSource{
				Type:   "card",
				Number: string(card.PAN),
				Balance: card.Balance,
			}
			sources = append(sources, source)
		}
	}
	return sources
	
	
}				
				

