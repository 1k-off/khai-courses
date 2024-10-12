package gen

import (
	"math/rand"
	"sort"
)

func Zeros(count int) []int {
	data := make([]int, count)
	return data
}

func Ones(count int) []int {
	data := make([]int, count)
	for i := range data {
		data[i] = 1
	}
	return data
}

func Ordered(count int, minValue int, order string) []int {
	data := make([]int, count)
	for i := 0; i < count; i++ {
		data[i] = minValue + i
	}

	if order == "desc" {
		sort.Sort(sort.Reverse(sort.IntSlice(data)))
	}

	return data
}

func Random(count int, minValue int, maxValue int) []int {
	data := make([]int, count)
	for i := range data {
		data[i] = rand.Intn(maxValue-minValue) + minValue
	}
	return data
}
