package card

import (
	"fmt"
	"lesson-5/pkg/bank/types"
)

// Test for successful withdrawal
func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20000, Currency: types.TJS, Active: true}, 10000)
	fmt.Println(result.Balance)
	// Output: 10000
}

// Test for insufficient funds
func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 0, Currency: types.TJS, Active: true}, 10000)
	fmt.Println(result.Balance)
	// Output: 0
}

// Test for inactive card
func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 20000, Currency: types.TJS, Active: false}, 10000)
	fmt.Println(result.Balance)
	// Output: 20000
}

// Test for limit 
func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 25000, Currency: types.TJS, Active: true}, 20_001)
	fmt.Println(result.Balance)
	// Output: 25000
}