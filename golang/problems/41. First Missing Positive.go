package problems

func FirstMissingPositive(nums []int) int {
	for n, _ := range nums {
		for nums[n] > 0 && nums[n] <= len(nums) && nums[n] != nums[nums[n] - 1] {
			nums[n], nums[nums[n] - 1]  = nums[nums[n] - 1], nums[n]
		}
	}
	for n, v := range nums {
		if v != n + 1{
			return n + 1
		}
	}
	return len(nums) + 1
}