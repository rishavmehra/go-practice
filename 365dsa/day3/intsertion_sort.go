package main

import "fmt"

func main() {
	arr := []int{23, 1, 10, 5, 2}
	res := IntsertionSort(arr)
	fmt.Println(res)
}

func IntsertionSort(input []int) []int {
	for i := 1; i < len(input); i++ {
		key := input[i]
		j := i - 1

		for j >= 0 && input[j] > key {
			input[j+1] = input[j]
			j--
		}
		input[j+1] = key
	}
	return input
}
