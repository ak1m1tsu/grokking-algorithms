package sort

func Selection(array []int) {
	if array == nil || len(array) == 0 {
		return
	}

	var minIdx int
	for i := 0; i < len(array); i++ {
		minIdx = i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}
		array[i], array[minIdx] = array[minIdx], array[i]
	}
}
