package main

import (
	"lesson-4/pkg/commission"
	"log"
)

func main() {
	
	result := commission.Calculate(9_999_999)
	log.Print(result)

}

// Commented out old code:
// wallet := 100.0
// card := 3120.5
// balance := caculateBalance(wallet, card)
// log.Print(balance)
//
// func caculateBalance(wallet float64, card float64) (balance float64) {
//	balance = wallet + card
//	return balance
// }
