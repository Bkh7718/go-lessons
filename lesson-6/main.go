package main

import (
	"fmt"
	"lesson-5/pkg/bank/card"
	"lesson-5/pkg/bank/types"
)

func main() {
	card := card.IssueCard(types.USD,  "white", "John Snow")
	fmt.Printf("Issued new card: %+v\n", card)
}
	