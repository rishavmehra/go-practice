package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 7, 1}

	for i := 0; i < len(arr)-1; i++ {
		MidIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[MidIndex] {
				MidIndex = j
			}
		}
		arr[i], arr[MidIndex] = arr[MidIndex], arr[i]

	}
	fmt.Println(arr)
}
