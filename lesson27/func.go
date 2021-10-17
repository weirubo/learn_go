package lesson27

func Sum(a, b int) int {
	return a + b
}

func Avg(arg ...int) float64 {
	sum := 0
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return float64(sum / len(arg))
}
