package problems
// The main idea is the same with problem Linked List Cycle II,https://leetcode.com/problems/linked-list-cycle-ii/.
// Use two pointers the fast and the slow.
// The fast one goes forward two steps each time, while the slow one goes only step each time.
// They must meet the same item when slow==fast.
// In fact, they meet in a circle, the duplicate number must be the entry point of the circle when visiting the array from nums[0].
// Next we just need to find the entry point.
// We use a point(we can use the fast one before) to visit form begining with one step each time, do the same job to slow. When fast==slow, they meet at the entry point of the circle.
// time o(n)	space o(1)

func FindDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}