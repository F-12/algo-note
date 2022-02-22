/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	nmap := make(map[int]int, len(nums))
	for idx, n := range nums {
		left := target - n
		if lidx, ok := nmap[left]; ok {
			return []int{lidx, idx}
		}
		nmap[n] = idx
	}
	return nil
}

// @lc code=end

