/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] == val {
			r++
			continue
		}
		nums[l] = nums[r]
		l++
		r++
	}
	return l
}

// @lc code=end

