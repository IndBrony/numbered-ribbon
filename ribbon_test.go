package ribbon

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	sort(arr)
	assert.Equal(t, true, reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5}))
	arr = []int{2, 4, 5, 1, 3}
	sort(arr)
	assert.Equal(t, true, reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5}))
}

func TestTotalPensTravelledDistance(t *testing.T) {
	basicTest := func(t *testing.T, nums []int, distancePerElement, expected int) {
		output := TotalPensTravelledDistance(nums, distancePerElement)
		assert.Equal(t, expected, output)
	}
	basicTest(t, []int{1, 2, 3, 4, 5}, 10, 40)
	basicTest(t, []int{40, 50, 20, 21, 22, 23, 24, 25}, 10, 150)
	basicTest(t, []int{100, 95, 90, 80, 40, 10, 8}, 10, 120)
}
func TestSolution(t *testing.T) {
	basicTest := func(t *testing.T, nums string, expected int, isError bool) {
		output, err := Solution(nums)
		if isError {
			assert.NotEqual(t, nil, err)
		} else {
			assert.Equal(t, nil, err)
		}
		assert.Equal(t, expected, output)
	}
	basicTest(t, "1 2 3 4 5", 40, false)
	basicTest(t, "40 50 20 21 22 23 24 25", 150, false)
	basicTest(t, "100 95 90 80 40 10 8", 120, false)

	basicTest(t, "a 95 90 80 40 10 8", -1, true)
}
