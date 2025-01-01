package array

func EqualSubArrays(list []int) [][]int {
	output := make([][]int, 0)
	if len(list) < 2 {
		return output
	}

	splitPoint := findSplitPoint(list)
	if splitPoint == -1 || splitPoint == len(list) {
		return output
	}

	output = append(output, list[0:splitPoint])
	output = append(output, list[splitPoint:])
	return output

}

func findSplitPoint(list []int) int {
	lsum := 0
	for i := range len(list) {
		lsum = lsum + list[i]

		rsum := 0
		for j := i + 1; j < len(list); j++ {
			rsum = rsum + list[j]
		}

		if lsum == rsum {
			return i + 1
		}

	}
	return -1
}
