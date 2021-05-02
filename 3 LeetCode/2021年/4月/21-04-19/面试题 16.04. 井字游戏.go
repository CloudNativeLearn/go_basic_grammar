package main

// 设计一个算法，判断玩家是否赢了井字游戏。输入是一个 N x N 的数组棋盘，由字符" "，"X"和"O"组成，其中字符" "代表一个空位。
//
//以下是井字游戏的规则：
//
//玩家轮流将字符放入空位（" "）中。
//第一个玩家总是放字符"O"，且第二个玩家总是放字符"X"。
//"X"和"O"只允许放置在空位中，不允许对已放有字符的位置进行填充。
//当有N个相同（且非空）的字符填充任何行、列或对角线时，游戏结束，对应该字符的玩家获胜。
//当所有位置非空时，也算为游戏结束。
//如果游戏结束，玩家不允许再放置字符。
//如果游戏存在获胜者，就返回该游戏的获胜者使用的字符（"X"或"O"）；如果游戏以平局结束，则返回 "Draw"；如果仍会有行动（游戏未结束），则返回 "Pending"。
//
//示例 1：
//
//输入： board = ["O X"," XO","X O"]
//输出： "X"
//示例 2：
//
//输入： board = ["OOX","XXO","OXO"]
//输出： "Draw"
//解释： 没有玩家获胜且不存在空位
//示例 3：
//
//输入： board = ["OOX","XXO","OX "]
//输出： "Pending"
//解释： 没有玩家获胜且仍存在空位
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/tic-tac-toe-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	tictactoe([]string{" O    X", " O     ", "     O ", "XXXXXXX", "  O    ", " O   O ", "O  X OO"})
}

func tictactoe(board []string) string {
	whetherHaveKong := false //是否含有空格

	// 判断横
	// 判断第一行
	yiyang := true
	for i := 1; i < len(board[0]); i++ {
		if board[0][i] == ' ' {
			whetherHaveKong = true
		}
		if board[0][i-1] != board[0][i] {
			yiyang = false
		}
	}
	if yiyang  && board[0][0]!=' '{
		return string(board[0][0])
	}

	left := true
	right := true
	for i := 1; i < len(board); i++ {
		// 左对角判断
		if left == true && board[i][i] != board[i-1][i-1] {
			left = false
		}
		// 右对角判断
		if right == true && board[i][len(board)-i-1] != board[i-1][len(board)-i] {
			right = false
		}
		if right == false && left == false {
			break
		}
	}
	if left && board[0][0]!=' ' {
		return string(board[0][0])
	}
	if right &&board[0][len(board[0])-1]!=' ' {
		return string(board[0][len(board[0])-1])
	}

	// 0 1 2
	// 0 1 2
	// 0 1 2

	// 判断其他行
	for i := 1; i < len(board); i++ {
		yiyang = true
		for j := 0; j < len(board[i]); j++ {
			if j != 0 && board[i][j-1] != board[i][j] {
				yiyang = false
			}

		}

		if yiyang && board[i][0]!=' ' {
			return string(board[i][0])
		}

		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == ' ' {
				whetherHaveKong = true
			}
			// 判断与上方是否合适
			if board[i-1][j] == '0' {
				buyes := []byte(board[i])
				buyes[j] = '0'
				board[i] = string(buyes)
			} else if board[i-1][j] != board[i][j] {
				buyes := []byte(board[i])
				buyes[j] = '0'
				board[i] = string(buyes)
			}
		}

	}

	// 判断列
	for i := 0; i < len(board[len(board)-1]); i++ {
		if board[len(board)-1][i] == 'O' || board[len(board)-1][i] == 'X' {
			return string(board[len(board)-1][i])
		}
	}

	if !whetherHaveKong {
		return "Draw"
	} else {
		return "Pending"
	}

}
