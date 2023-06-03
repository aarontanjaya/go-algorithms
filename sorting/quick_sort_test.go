package sorting_test

import (
	"testing"

	"github.com/aarontanjaya/go-algorithms/sorting"
)

func TestQuickSortInts(t *testing.T) {
	t.Run("Should result in correct order for non reversed integer case", func(t *testing.T) {
		for j, arr := range IntCases {
			sorting.QuickSortInts(arr, false)
			for i := 0; i < len(arr)-2; i++ {
				if arr[i] > arr[i+1] {
					t.Errorf("Error on integer testcase %d index %d, %d expected to be smaller than %d", j, i, arr[i], arr[i+1])
				}
			}
		}
	})

	t.Run("Should result in correct order for reversed integer case", func(t *testing.T) {
		for j, arr := range IntCases {
			sorting.QuickSortInts(arr, true)
			for i := 0; i < len(arr)-2; i++ {
				if arr[i] < arr[i+1] {
					t.Errorf("Error on integer testcase %d index %d, %d expected to be greater than %d", j, i, arr[i], arr[i+1])
				}
			}
		}
	})
}
