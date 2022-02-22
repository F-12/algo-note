/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

import "sort"

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	combined := make([][]int, 0)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	return dfs(candidates, combined, make([]int, 0), 0, target)
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
		combined = dfs(candidates, combined, path, from+1, target-candidates[from])
		path = path[:len(path)-1]
		for ; from+1 < len(candidates) && candidates[from+1] == candidates[from]; from++ {

		}
	}

	return combined
}

// @lc code=end

