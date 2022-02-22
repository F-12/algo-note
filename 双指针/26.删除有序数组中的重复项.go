/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 * 双指针法
 * - 左指针指向有效数组最后下标
 * - 右指针指向遍历位置
 * - 循环终止条件：右指针遍历结束
 * - 左指针推进条件：左右指针值不相等，左指针填入右指针值
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	l, r := 0, 1
	for r < len(nums) {
		if nums[r] == nums[l] {
			r++
			continue
		}
		nums[l+1] = nums[r]
		l++
		r++
	}
	return l + 1
}

// @lc code=end

