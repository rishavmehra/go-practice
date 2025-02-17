package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{13, 46, 24, 52, 20, 9}
	got := BubbleSort(arr)
	want := []int{9, 13, 20, 24, 46, 52}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BubbleSort() = %v; want %v", got, want)
	}

}
