/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 */

// @lc code=start
func maxArea(height []int) int {
	// 双指针法
	area, l, r := 0, 0, len(height)-1
	// 左右指针重合循环结束
	for l < r {
		// 记录容器的高对应的下标
		h := 0
		if height[l] <= height[r] {
			h = l
			l++
		} else {
			h = r
			r--
		}
		narea := height[h] * (r - l + 1)
		// fmt.Printf("(%v,%v,%v,%v)\n", h, r, l, narea)
		if narea >= area {
			area = narea
		}
	}
	return area
}

// @lc code=end

