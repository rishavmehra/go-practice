package fibonaccisequence

func FibonacciRecur(num int) int {
	if num <= 1 {
		return num
	}
	if num == 2 {
		return 1
	}

	return FibonacciRecur(num-1) + FibonacciRecur(num-2)
}

func Fibonaccisequence(n int) []int {
	var res = []int{}
	for i := 0; i <= n; i++ {
		result := fibonacci(i)
		res = append(res, result)
	}
	return res
}

func fibonacci(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
	}
	return a
}
