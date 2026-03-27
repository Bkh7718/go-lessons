package payment

import (
	"go-lessons/pkg/bank/types"
)

// Max takes a slice of payments and returns the payment with the maximum amount. It iterates through the payments, comparing each payment's amount to find the maximum.
func Max(payments []types.Payment) types.Payment {

	max := payments[0]

	for _, payment := range payments {

		if payment.Amount > max.Amount {

			max = payment
		}

	}
	payments[0] = max //return max
	return payments[0]
}
