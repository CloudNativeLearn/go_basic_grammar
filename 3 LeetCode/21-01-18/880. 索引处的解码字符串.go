package main

import "fmt"

// 给定一个编码字符串 S。请你找出 解码字符串 并将其写入磁带。解码时，从编码字符串中 每次读取一个字符 ，并采取以下步骤：
//
//如果所读的字符是字母，则将该字母写在磁带上。
//如果所读的字符是数字（例如 d），则整个当前磁带总共会被重复写d-1 次。
//现在，对于给定的编码字符串 S 和索引 K，查找并返回解码字符串中的第K个字母。

// 示例 1：
//
//输入：S = "leet2code3", K = 10
//输出："o"
//解释：
//解码后的字符串为 "leetleetcodeleetleetcodeleetleetcode"。
//字符串中的第 10 个字母是 "o"。

//示例 2：
//
//输入：S = "ha22", K = 5
//输出："h"
//解释：
//解码后的字符串为 "hahahaha"。第 5 个字母是 "h"。

//示例 3：
//
//输入：S = "a2345678999999999999999", K = 1
//输出："a"
//解释：
//解码后的字符串为 "a" 重复 8301530446056247680 次。第 1 个字母是 "a"。

func main() {
	x := "xx2q"
	fmt.Println(GetTHeStrLength(&x))
	fmt.Println(decodeAtIndex("ha22",5))
}

func decodeAtIndex(S string, K int) string {
	sNO := GetTHeStrLength(&S)
	K = K-1
	for i:=len(S)-1;i>=0;i-- {
		if 48<S[i] && S[i]<58 {
			// 表示这个字符为数字
			sNO = sNO/int(S[i]-48)
			K = K%sNO
		}else {
			sNO = sNO -1
			if K==sNO{
				return string(S[i])
			}
		}
	}


	return string(S[0])

}

// 获得字符串的长度
func GetTHeStrLength(x *string) int {
	result := 0
	mid := 0
	y := *x
	for i := 0; i < len(y); i++ {
		if 48 < y[i] && y[i] < 58 {
			// 表示当前元素为 0-10
			result = (result + mid) * (int(y[i]) - 48)
			mid = 0
		} else {
			// 表示当前元素为字母
			mid = mid + 1
		}
	}
	result = result + mid
	return result
}
