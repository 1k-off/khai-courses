package bubble

func Sort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}
