package transfer

func Total(amount float64,) float64 {
	bonusAmount := amount * 0.005
	return bonusAmount + amount
}