package ribbon

import (
	"math"
	"strconv"
	"strings"
)

//Solution is (as the name imply) the solution to Jum'at Hek challenge
//it takes a string of numbers separated by a space as input
//and outputs the total distance travelled by pen (and an error if any)
func Solution(input string) (int, error) {
	//split the input and convert it to int slice
	numsString := strings.Split(input, " ")
	nums := []int{}
	for _, numString := range numsString {
		num, err := strconv.Atoi(numString)
		if err != nil {
			return -1, err
		}
		nums = append(nums, num)
	}
	//calculate the total distance travelled with distance 10 between each element
	return TotalPensTravelledDistance(nums, 10), nil
}

func TotalPensTravelledDistance(nums []int, distancePerElement int) (totalDistance int) {
	currPenPosition := 0
	//map the (value -> index)
	mapOfIndex := make(map[int]int)
	for index, value := range nums {
		mapOfIndex[value] = index
	}

	//sort the input and store it on separate slice
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort(sorted)

	//calculate the distance between current and the next pen position
	//next pen position is the lowest value on the ribbon and then ascending
	for _, value := range sorted {
		nextPenPosition := mapOfIndex[value]
		distance := math.Abs(float64(nextPenPosition - currPenPosition))
		currPenPosition = nextPenPosition
		totalDistance += int(distance) * distancePerElement
	}
	return
}

func sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		k := i
		for i != 0 && arr[i] < arr[i-1] {
			temp := arr[i-1]
			arr[i-1] = arr[i]
			arr[i] = temp
			i--
		}
		i = k
	}
}
