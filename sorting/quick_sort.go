package sorting

import (
	"sort"
)

func partition(data sort.Interface, low int, high int) int {
	pivot := high
	i := low
	for j := low; j < high; j++ {
		if data.Less(j, pivot) {
			data.Swap(i, j)
			i++

		}
	}

	data.Swap(i, high)
	return i
}

func QuickSort(data sort.Interface, low int, high int) {
	if low < high {
		p := partition(data, low, high)

		QuickSort(data, p+1, high)
		QuickSort(data, low, p-1)
	}

}

func QuickSortInts(arr []int, reverse bool) {
	if reverse {
		QuickSort(sort.Reverse(sort.IntSlice(arr)), 0, len(arr)-1)
		return
	}
	QuickSort(sort.IntSlice(arr), 0, len(arr)-1)
}

func QuickSortStr(arr []string, reverse bool) {
	if reverse {
		QuickSort(sort.Reverse(sort.StringSlice(arr)), 0, len(arr)-1)
		return
	}
	QuickSort(sort.StringSlice(arr), 0, len(arr)-1)
}
