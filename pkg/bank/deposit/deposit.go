package deposit

// Calculate рассчитывает параметры вклада.
func Calculate(amount int, currency string) (min, max int) {

	minPercent, maxPercent := PercentFor(currency)

	min = amount * minPercent / 100
	max = amount * maxPercent / 100
	return min /12, max /12
}

func PercentFor(currency string) (min, max int)  {
	

switch currency {
case "TJS":
	return 4, 6

case "USD":
		return 3, 4
default:	
	return 0, 0
}
		

}