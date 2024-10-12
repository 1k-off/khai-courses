package radixsort

func Sort(arr []int) []int {
	m := maxValue(arr)
	for exp := 1; m/exp > 0; exp *= 10 {
		countSort(arr, exp)
	}
	return arr
}

func maxValue(arr []int) int {
	m := arr[0]
	for _, v := range arr {
		if v > m {
			m = v
		}
	}
	return m
}

func countSort(arr []int, exp int) []int {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	for i := 0; i < n; i++ {
		count[(arr[i]/exp)%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		output[count[(arr[i]/exp)%10]-1] = arr[i]
		count[(arr[i]/exp)%10]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
	return arr
}
