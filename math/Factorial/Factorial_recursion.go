package factorial

func FactorialRecusion(num int) int {
	if num == 1 {
		return 1
	}
	return num * FactorialRecusion(num-1)
}
