package linearsearch

func Linear_search(number []int, n int) int {

	for i := 0; i < len(number); i++ {
		if number[i] == n {
			return i
		}
	}
	return -1
}
