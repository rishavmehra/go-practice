package array

import "math"

func FindDuplicate(list []int) int {
	for _, item := range list {
		itemIndex := int(math.Abs(float64(item))) - 1
		if list[itemIndex] < 0 {
			return item
		}
		list[itemIndex] *= -1
	}
	return -1
}
