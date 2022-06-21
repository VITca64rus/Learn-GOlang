func M() float64 {
	return (p * v)
}

func W() float64 {
	m := M()
	return (math.Sqrt(k/m))
}

func T() float64 {
	w := W()
	return (6/w)
}