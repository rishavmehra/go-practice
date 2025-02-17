package main

import (
	"reflect"
	"testing"
)

func TestIntsertionSort(t *testing.T) {
	arr := []int{23, 1, 10, 5, 2}
	got := IntsertionSort(arr)
	want := []int{1, 2, 5, 10, 23}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got is %v, want is %v", got, want)
	}

}
