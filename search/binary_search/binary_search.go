package binarysearch

func Binary_search(numbers []int, n int) bool {

	low := 0
	high := len(numbers)

	for low < high {
		mid := low + (high-low)/2
		value := numbers[mid]

		if value == n {
			return true
		} else if value > n {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return false
}
