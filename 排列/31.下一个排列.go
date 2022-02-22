/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 * - 从右开始行进到单调递增结束位置 j, a[j] < a[j+1]
 * - 从a[j+1]到a[n-1]找到大于a[j]的最小整数a[i]
 * - a[i]放到j处
 * - 剩余数字递增放到 j+1到n-1位置(只需要逆序即可，[j+1,n)区间是单调递减的)
 */

// @lc code=start
func nextPermutation(nums []int) {
	j := len(nums) - 2
	for j >= 0 && nums[j] >= nums[j+1] {
		j--
	}
	if j >= 0 {
		i := j
		for i < len(nums)-1 && nums[i+1] > nums[j] {
			i++
		}
		swap(nums, i, j)
	}
	lo, hi := j+1, len(nums)-1
	for lo < hi {
		swap(nums, lo, hi)
		lo++
		hi--
	}
}

func swap(nums []int, i, j int) {
	v := nums[i]
	nums[i] = nums[j]
	nums[j] = v
}

// @lc code=end

