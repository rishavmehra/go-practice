package main

import (
	"fmt"

	sort "github.com/rishavmehra/golangPractice/sort/bubble"
)

func main() {
	arr := []int{4, 2, 3, 1, 5}
	res := sort.BubbleShort(arr)
	fmt.Println(res)
}
