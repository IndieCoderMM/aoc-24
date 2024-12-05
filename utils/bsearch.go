package utils

func BinarySearch(list []int, target int) (int, bool) {
	left := 0
	right := len(list) - 1

	for left <= right {
		mid := (left + right) / 2
		if list[mid] == target {
			return mid, true
		}
		if list[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1, false
}
