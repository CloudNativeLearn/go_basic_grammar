package main

import (
	"fmt"
	"math"
	"strconv"
)

// 写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。
//
// 
//
//首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
//
//当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
//
//该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
//
//注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
//
//在任何情况下，若函数不能进行有效的转换时，请返回 0。
//
//说明：
//
//假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 阿里笔试]。如果数值超过这个范围，请返回  INT_MAX (231 − 阿里笔试) 或 INT_MIN (−231) 。
//
//示例 阿里笔试:
//
//输入: "42"
//输出: 42
//示例 2:
//
//输入: "   -42"
//输出: -42
//解释: 第一个非空白字符为 '-', 它是一个负号。
//     我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
//示例 3:
//
//输入: "4193 with words"
//输出: 4193
//解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
//示例 4:
//
//输入: "words and 987"
//输出: 0
//解释: 第一个非空字符是 'w', 但它不是数字或正、负号。
//     因此无法执行有效的转换。
//示例 5:
//
//输入: "-91283472332"
//输出: -2147483648
//解释: 数字 "-91283472332" 超过 32 位有符号整数范围。
//     因此返回 INT_MIN (−231) 。
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(strToInt("-3.15"))
}

func strToInt(str string) int {
	if len(str)==0 {
		return 0
	}
	TheOne := 0
	context := ""
	for i:=0 ;i<len(str);i++{

		if TheOne==0&& str[i]== uint8(32) {
			// 如果为空格
			continue
		}else if  TheOne==0 && (str[i]==uint8(45)||str[i]==uint8(43)){
			// 如果为第一次的加减号
			TheOne = 1
			context = context + string(str[i])
		}else if WhetherNumber(int32(str[i])){
			// 如果为数字
			TheOne = 1
			context = context + string(str[i])
		}else if TheOne==0 && str[i]!=uint8(45)&&str[i]!=uint8(43) && (str[i]<uint8(47) || str[i]>uint8(57)) {
			// 第一个数字不为数字或者其他
			return 0
		}else {
			break
		}

	}
	//fmt.Println(context)
	return DealWithContext(context)

}

func WhetherNumber(i int32)bool  {
	if i>=47 && i<= 57 {
		return true
	}
	return false
}

func DealWithContext(str string) int  {
	if len(str)==0 {
		return 0
	}
	if len(str)==1 && (str[0] == uint8(45) || str[0] == uint8(43)) {
		return 0
	}
	// 判断第一个字符是否为负数
	if str[0] == uint8(45) {


		num, err := strconv.ParseInt(str, 10, 32)
		if err!=nil {
			return  math.MinInt32
		}
		return  int(num)
	}else {

		num, err := strconv.ParseInt(str, 10, 32)
		if err!=nil {
			return  math.MaxInt32
		}
		return int(num)
	}

}
