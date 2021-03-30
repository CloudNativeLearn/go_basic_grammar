package main

//
//字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
//
//
//
//示例 1:
//
//输入:
//first = "pale"
//second = "ple"
//输出: True
//
//
//示例 2:
//
//输入:
//first = "pales"
//second = "pal"
//输出: False
//
//
//func main() {
//	fmt.Println(oneEditAway("abcdxabcde", "abcdeabcdx"))
//}

func oneEditAway(first string, second string) bool {
	Lfirst := len(first)
	LSecond := len(second)
	if Lfirst == 0 && LSecond == 0 {
		return true
	} else if (LSecond == 0 && Lfirst == 1) || (LSecond == 1 && Lfirst == 0) {
		return true
	} else if Lfirst == 1 && LSecond == 1 {
		return true
	}

	l := 0
	r := 0
	Number := 0
	// 默认把第一个变成第二个
	for l < Lfirst || r < LSecond {
		if Number > 1 {
			return false
		}

		if r >= LSecond || l >= Lfirst {
			if Number = Number + 1; Number > 1 {
				return false
			}
			l = l + 1
			r = r + 1
			continue
		}

		if first[l] == second[r] {
			l = l + 1
			r = r + 1
			continue
		} else {

			if l+1 < Lfirst && r+1 < LSecond {
				if first[l] == second[r+1] {
					r = r + 1
				} else if first[l+1] == second[r] {
					l = l + 1
				} else if first[l+1] == second[r+1] {
					l = l + 1
					r = r + 1
				} else {
					return false
				}
			} else if l+1 < Lfirst {
				l = l + 1
			} else if r+1 < LSecond {
				r = r + 1
			} else {
				l = l + 1
				r = r + 1
			}

			if Number = Number +1;Number>1{
				return false
			}

		}
	}
	return true
}
