package payment

import (
	"go-lessons/pkg/bank/types"
	"fmt"
)

// Max takes a slice of payments and returns the payment with the maximum amount. It iterates through the payments, comparing each payment's amount to find the maximum.
func ExampleMax() {
	payments := []types.Payment{
		{ID: 1, Amount: 1000},
		{ID: 2, Amount: 5000},
		{ID: 3, Amount: 3000},
		{ID: 4, Amount: 2000},
		{ID: 5, Amount: 4000},
	}
	max := Max(payments)
	fmt.Println(max.ID, max.Amount)
	// Output: 2 5000
}