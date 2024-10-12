package gostd

import "sort"

func Sort(arr []int) []int {
	sort.Ints(arr)
	return arr
}
