package main

import (
	"fmt"

	"github.com/rishavmehra/golangPractice/search/questions"
)

func main() {
	// linearSearch := linearsearch.Linear_search([]int{1, 2, 3, 4, 5, 6}, 6)
	// binarysearch := binarysearch.Binary_search([]int{1, 2, 3, 4, 5, 6}, 7)
	// fmt.Println(binarysearch)

	breaks1 := []bool{false, false, true, false, true, false, false}
	fmt.Printf("Result: %v\n", questions.TwoCrystalBalls(breaks1))

}
