package search

func Binary(sortedArray []int, target int) int {
	low, high := 0, len(sortedArray)-1

	for low <= high {
		mid := (low + high) / 2
		guess := sortedArray[mid]
		if guess == target {
			return mid
		}
		if guess > target {
			high = mid - 1
		}
		if guess < target {
			low = mid + 1
		}
	}

	return -1
}
