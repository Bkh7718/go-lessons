package card

import (
	"fmt"
	"go-lessons/pkg/bank/types"
)

// Test for successful withdrawal
func ExampleWithdraw_positive() {
	card := types.Card{Balance: 20000, Currency: types.TJS, Active: true}
	Withdraw(&card, 10000)
	fmt.Println(card.Balance)
	// Output: 10000
}

// Test for insufficient funds
func ExampleWithdraw_noMoney() {
	card := types.Card{Balance: 0, Currency: types.TJS, Active: true}
	Withdraw(&card, 10000)
	fmt.Println(card.Balance)
	// Output: 0
}

// Test for inactive card
func ExampleWithdraw_inactive() {
	card := types.Card{Balance: 20000, Currency: types.TJS, Active: false}
	Withdraw(&card, 10000)
	fmt.Println(card.Balance)
	// Output: 20000
}

// Test for limit
func ExampleWithdraw_limit() {
	card := types.Card{Balance: 25000, Currency: types.TJS, Active: true}
	Withdraw(&card, 20_001)
	fmt.Println(card.Balance)
	// Output: 25000
}

func ExampleDeposit_positive() {
	card := types.Card{Balance: 10000, Currency: types.TJS, Active: true}
	Deposit(&card, 20000)
	fmt.Println(card.Balance)
	// Output: 30000
}

func ExampleDeposit_inactive() {
	card := types.Card{Balance: 10000, Currency: types.TJS, Active: false}
	Deposit(&card, 20000)
	fmt.Println(card.Balance)
	// Output: 10000
}

func ExampleDeposit_limit() {
	card := types.Card{Balance: 10000, Currency: types.TJS, Active: true}
	Deposit(&card, 50_001)
	fmt.Println(card.Balance)
	// Output: 10000
}

func ExampleAddBonus_positive() {
	card := types.Card{Balance: 10000, Currency: types.TJS, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 10024
}

func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 10000, Currency: types.TJS, Active: false}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 10000
}

func ExampleAddBonus_negativeMoney() {
	card := types.Card{Balance: -10000, Currency: types.TJS, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: -10000
}

func ExampleAddBonus_limit() {
	card := types.Card{Balance: 10_000_000, Currency: types.TJS, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 10005000
}
