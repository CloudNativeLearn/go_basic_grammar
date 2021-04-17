package main

// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
// 
//
//示例 1：
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//示例 2：
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//示例 3：
//
//输入：s = ""
//输出：0
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-valid-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//func main() {
//	println(longestValidParentheses("(()(((()"))
//}

func longestValidParentheses(s string) int {
	if len(s)==0 {
		return 0
	}
	stark := []int32{}
	Result := 0
	ResultList := []int{}
	zhongduan := false
	//_ = zhongduan
	for _, v := range s {
		switch v {
		case int32(40): // (
			if len(stark)==0 {
				ResultList = append(ResultList, Result)
				Result = 0
				if zhongduan {
					zhongduan = false
				}
				stark = append(stark, v)
			}else {
				stark = append(stark, v)
			}

		case int32(41): // )
			if len(stark) == 0 {
				// 发生中断 讲result append进入list   result变为空
				zhongduan = true
				ResultList = append(ResultList, Result)
				Result = 0

			} else {
				stark = stark[:len(stark)-1]
				Result = Result + 2
				if len(stark)==0 {
					if !zhongduan && len(ResultList)>0 {
						Result = Result + ResultList[len(ResultList)-1]
					}
				}
			}

		}
	}
	ResultList = append(ResultList, Result)

	ThMax :=0
	for _,v := range ResultList{
		if ThMax<v {
			ThMax = v
		}
	}
	return ThMax
}

// 判断是否有数据 有数据返回true 并且删除元素
func DealStart(stark *[]int32) (*[]int32, bool) {
	if len(*stark) == 0 {
		return stark, false
	} else {
		mid := (*stark)[:len(*stark)-1]
		return &mid, true
	}

}
