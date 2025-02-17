package main

import "fmt"

func main() {
	arr := []int{23, 1, 10, 5, 2}
	res := MergeSort(arr)
	fmt.Println(res)
}

func MergeSort(input []int) []int {
	if len(input) < 2 {
		return input
	}
	left := MergeSort(input[:len(input)/2])
	right := MergeSort(input[len(input)/2:])
	return merge(left, right)
}

func merge(a, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
