package src

type fnCalculate func(x, y int) int

func tax(x, y int, cb fnCalculate) float64 {
	k := 0.10
	r := float64(cb(x, y))
	return r * k
}

func Calculate(x, y int, option string) float64 {
	mp := map[string]fnCalculate{
		"sum": func(x, y int) int {
			return x + y
		},
		"sub": func(x, y int) int {
			return x - y
		},
	}

	fn, exists := mp[option]

	if !exists {
		panic("Option not found")
	}

	return tax(x, y, fn)
}
