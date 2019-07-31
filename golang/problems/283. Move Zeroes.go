package problems

// time o(n)	space o(1)

func moveZeroes(nums []int)  {
	nz := 0
	for i := 0; i <= len(nums) - 1; i++ {
		if nums[i] != 0 {
			nums[nz] = nums[i]
			nz++
		}
	}
	for i := nz; i <= len(nums) - 1; i++ {
		nums[i] = 0
	}
}