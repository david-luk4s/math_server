package math

func Add(a, b int64) int64 {
	return a + b
}

func Subtract(a, b int64) int64 {
	return a - b
}

func Divide(a, b int64) float64 {
	if b == 0 {
		return float64(0)
	}
	return float64(a / b)
}

func Multiply(a, b int64) int64 {
	return a * b
}
