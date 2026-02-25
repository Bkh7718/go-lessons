package commission

func Calculate(MonthlyAmount float64) float64 {

	commissionRate := 0.005
	commission := MonthlyAmount * commissionRate
	return (commission)
}
