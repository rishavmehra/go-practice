package array

func RotateKSteps(list []int, k int) {
	replace(list, 0, len(list)-k-1)
	replace(list, 0, len(list)-1)
	replace(list, 0, k-1)
}

func replace(list []int, start int, end int) {
	for i := start; i <= start+end/2 && i < end-i+start; i++ {
		list[i], list[end-i+start] = list[end-i+start], list[i]
	}
}
