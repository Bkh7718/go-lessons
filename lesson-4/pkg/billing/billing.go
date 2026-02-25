package billing

func Calculate1000(Trafic float64) float64 {
	abonPlata := 3500.0
	if Trafic > 1000 {
		abonPlata = 3500.0 + (Trafic-1000)*6
	}
	return abonPlata
}

func Calculate2000(Trafic float64) float64 {
	abonPlata := 5500.0
	if Trafic > 2000 {
		abonPlata = 5500.0 + (Trafic-2000)*6
	}
	return abonPlata
}

func Calculate3000(Trafic float64) float64 {
	abonPlata := 7000.0
	if Trafic > 3000 {
		abonPlata = 7000.0 + (Trafic-3000)*6
	}
	return abonPlata
}

func Calculate5000(Trafic float64) float64 {
	abonPlata := 9500.0
	if Trafic > 5000 {
		abonPlata = 9500.0 + (Trafic-5000)*6
	}
	return abonPlata
}
func Calculate10000(Trafic float64) float64 {
	abonPlata := 17000.0
	if Trafic > 10000 {
		abonPlata = 17000.0 + (Trafic-10000)*6
	}
	return abonPlata
}