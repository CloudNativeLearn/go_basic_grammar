package main

import "fmt"

//给出一个字符串 s（仅含有小写英文字母和括号）。
//
//请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。
//
//注意，您的结果中 不应 包含任何括号。
//
// 
//
//示例 1：
//
//输入：s = "(abcd)"
//输出："dcba"
//示例 2：
//
//输入：s = "(u(love)i)"
//输出："iloveu"
//示例 3：
//
//输入：s = "(ed(et(oc))el)"
//输出："leetcode"
//示例 4：
//
//输入：s = "a(bcdefghijkl(mno)p)q"
//输出："apmnolkjihgfedcbq"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(reverseParentheses("(abcd)"))
}

func reverseParentheses(s string) string {
	Stark2 := Stark{
		[]uint8{},
	}
	for i:=0;i<len(s);i++ {
		if string(s[i]) ==")"{
			temp:=[]uint8{}
			x := uint8(0)
			for string(x)!="("{
				if x!=0{
					temp = append(temp, x)
				}
				x= Stark2.POP()
			}
			Stark2.Info = append(Stark2.Info, temp...)
		}else {
			Stark2.add(s[i])
		}
	}

	result := ""
	for i:=0;i<len(Stark2.Info);i++{
		result = result+string(Stark2.Info[i])
	}
	return result
}

type Stark struct {
	Info []uint8
}

func (this *Stark)add(a uint8)  {
	this.Info = append(this.Info, a)
}

func (this *Stark)POP() uint8  {
	x := this.Info[len(this.Info)-1]
	this.Info = this.Info[:len(this.Info)-1]
	return x
}

func (this *Stark)GetLine()uint8  {
	return this.Info[len(this.Info)-1]
}
