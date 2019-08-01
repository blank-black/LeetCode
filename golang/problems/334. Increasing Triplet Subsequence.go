package problems
import "math"

// start with two largest values, as soon as we find a number bigger than both, while both have been updated, return true.
func IncreasingTriplet(nums []int) bool {
	small := math.MaxInt32
	big := math.MaxInt32
	for _, v := range nums {
		if v <= small {
			small = v
		} else if v <= big {
			big = v
		} else {
			return true
		}
	}
	return false
}