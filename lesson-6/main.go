package main

import (
	"fmt"
	"lesson-5/pkg/bank/card"
	"lesson-5/pkg/bank/types"
)

func main() {
	NewCard := card.IssueCard(types.USD,  "white", "John Snow")
	fmt.Printf("Issued new card: %+v\n", NewCard)
	// Attempt to withdraw from the card
	updatedCard := card.Withdraw(NewCard, 1500)
	fmt.Printf("Balance after withdrawal: %+v\n", updatedCard.Balance)
	
	
}
	