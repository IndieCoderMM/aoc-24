package utils

func QuickSort(arr []int, low int, high int, reverse bool) {
	if low < high {
		pivot := partition(arr, low, high, reverse)

		QuickSort(arr, low, pivot-1, reverse)
		QuickSort(arr, pivot+1, high, reverse)
	}
}

func partition(arr []int, low int, high int, reverse bool) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		condition := arr[j] <= pivot // Asc
		if reverse {
			condition = arr[j] > pivot // Desc
		}
		if condition {
			i += 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}
