package main

import (
	"fmt"
	"lesson-5/pkg/bank/deposit"
	"lesson-5/pkg/bank/transfer"
)

func main() {
	min, max := deposit.Calculate(1000, "USD")
	fmt.Println("Deposit profit:", min, max)
	fmt.Println("Transfer total:", transfer.Total(5_000))
}
