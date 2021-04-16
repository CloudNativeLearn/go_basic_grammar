package main

import "fmt"

// 括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。
//
//说明：解集不能包含重复的子集。
//
//例如，给出 n = 3，生成结果为：
//
//[
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/bracket-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(generateParenthesis(3))
}


func generateParenthesis(n int) []string {
result := []string{}
Stark(0,0,n,"",&result)
return result
}

func Stark(left int,right int,n int,MindString string,result *[]string)  {

	if right>left || right>n || left>n{
		return
	}

	if left+right==2*n && left==n && right==n{
		*result = append(*result, MindString)
		return
	}


	// 增加一个左括号
	Stark(left+1,right,n,MindString+"(",result)
	// 增加一个右括号
	Stark(left,right+1,n,MindString+")",result)
}


