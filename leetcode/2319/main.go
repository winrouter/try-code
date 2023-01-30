package main

import "fmt"

/*
如果一个正方形矩阵满足下述 全部 条件，则称之为一个 X 矩阵 ：

矩阵对角线上的所有元素都 不是 0
矩阵中所有其他元素都是 0
给你一个大小为 n x n 的二维整数数组 grid ，表示一个正方形矩阵。如果 grid 是一个 X 矩阵 ，返回 true ；否则，返回 false 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/check-if-matrix-is-x-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/



func checkXMatrix(grid [][]int) bool {
	n := len(grid[0])
	fmt.Println(n)

	for i := 0; i < n; i++ {
		if grid[i][i] == 0 {
			return false
		}
		grid[i][i] = 0

		if i == n-1-i {
			continue
		}

		if grid[i][n-1-i] == 0 {
			return false
		}
		grid[i][n-1-i] = 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	grid := [][]int {{5, 0, 20}, {0,5,0}, {6,0,2}}

	fmt.Println(checkXMatrix(grid))

}