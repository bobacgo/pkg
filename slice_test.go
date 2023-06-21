package main

import (
	"fmt"
	"testing"
)

func TestCrossMergeSlice(t *testing.T) {
	arr1 := []int{1, 2, 9}
	arr2 := []int{3, 7}
	arr3 := []int{5, 4, 3, 5, 7, 8}
	arr4 := []int{6, 9, 1}
	arr := CrossMergeSlice(arr1, arr2, arr3, arr4)
	fmt.Println(arr)
}
