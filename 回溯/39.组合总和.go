/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和

 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	return dfs(candidates, res, make([]int, 0), 0, target)
}

func dfs(candidates []int, combined [][]int, path []int, idx, target int) [][]int {
	if target == 0 {
		combined = append(combined, append([]int(nil), path...))
		return combined
	}

	if target < 0 || idx >= len(candidates) {
		return combined
	}
	for from := idx; from < len(candidates); from++ {
		// 选第idx个
		path = append(path, candidates[from])
		combined = dfs(candidates, combined, path, from, target-candidates[from])
		path = path[:len(path)-1]
	}

	return combined
}

// @lc code=end

