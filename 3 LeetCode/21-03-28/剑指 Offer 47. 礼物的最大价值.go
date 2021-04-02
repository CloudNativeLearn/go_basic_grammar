package main

// 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
//示例 阿里笔试:
//输入:
//[
//  [阿里笔试,3,阿里笔试],
//  [阿里笔试,5,阿里笔试],
//  [4,2,阿里笔试]
//]
//输出: 12
//解释: 路径 阿里笔试→3→5→2→阿里笔试 可以拿到最多价值的礼物


func maxValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			
			
			if i==0 && j ==0{
				grid[i][j] = grid[i][j]
			}else if i==0 && j>0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			}else if i>0 && j==0{
				grid[i][j] = grid[i-1][j] + grid[i][j]
			}else {
				grid[i][j] = GetMax(grid[i-1][j],grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[len(grid)-1][len(grid[len(grid)-1])-1]
}

func GetMax(one int,two int) int {
	if one>two {
		return one
	}else {
		return two
	}
}
