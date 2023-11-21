package main

func SumInts(m map[string]int64) int64 {
	sum := int64(0)
	for _, v := range m {
		sum += v
	}
	return sum
}

func SumFloats(m map[string]float64) float64 {
	sum := 0.0
	for _, v := range m {
		sum += v
	}
	return sum
}
