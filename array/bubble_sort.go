package array

func BubbleSort(input []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(input); i++ {
			if input[i] < input[i-1] {
				input[i-1], input[i] = input[i], input[i-1]
				swapped = true
			}
		}
	}
}
