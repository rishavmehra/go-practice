package main

import "fmt"

func main() {
	arr := []int{13, 46, 24, 52, 20, 9}
	res := BubbleSort(arr)
	fmt.Printf("%v", res)
}

// {13,46,24,52,20,9}
func BubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}
