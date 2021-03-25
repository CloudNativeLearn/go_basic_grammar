package main

import "fmt"

// 有重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合。
//
//示例1:
//
// 输入：S = "qqe"
// 输出：["eqq","qeq","qqe"]
//示例2:
//
// 输入：S = "ab"
// 输出：["ab", "ba"]


func main() {
	fmt.Println(permutation(""))
}


func permutation(S string) []string {

	
	mp := DealMap(S)

	result := []string{}
	recursiveString(mp,&result,"")
	
	return result
	
}
// qqe

// 递归处理字符串
func recursiveString(mp map[string]int,StList *[]string,context string)  {
	// 不全为0 有剩余

	if JudgeMap(mp) {
		// 遍历map 继续合并字符串
		for k,v := range mp{
			if v!=0 {
				c2 := context + k
				mp3 := DeepCopyMap(mp)
				mp3[k] = mp3[k]-1
				recursiveString(mp3,StList,c2)
			}
		}
	}else {
		// 无剩余  将字符串转递进去
		*StList = append(*StList, context)
	}


}


// 处理字符串 统计字符串中的字母个数
func DealMap(S string)map[string]int  {
	mp := map[string]int{}
	for _,v := range S{
		mp[string(v)] = mp[string(v)] +1
	}
	return mp
}

// 判断map中的  是否全部为0  如果为0返回false  如果有不为0返回true
func JudgeMap(m map[string]int) bool {
	for _,v :=range m{
		if v!=0 {
			return true
		}
	}
	return false
}

// map的深度拷贝
func DeepCopyMap(mp map[string]int)map[string]int  {
	newMap := map[string]int{}
	for k,v := range mp{
		newMap[k] = v
	}
	return newMap
}