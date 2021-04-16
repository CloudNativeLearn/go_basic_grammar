package main

import "fmt"

// 编写一种方法，对字符串数组进行排序，将所有变位词组合在一起。变位词是指字母相同，但排列不同的字符串。
//
//注意：本题相对原题稍作修改
//
//示例:
//
//输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//说明：
//
//所有输入均为小写字母。
//不考虑答案输出的顺序。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/group-anagrams-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}


func groupAnagrams(strs []string) [][]string {
	ResultMap := map[string][]string{}

	for _, V := range strs {
		newString := ""
		for _,v2 := range V{
			if len(newString)==0{
				newString = newString+string(v2)
			}else {
				i := 0
				for i<len(newString){
					if newString[i] > uint8(v2){
						break
					}
					i = i+1
				}

				newString = splitStringbyIndex(newString,i,string(v2))
			}
		}
		ResultMap[newString] = append(ResultMap[newString], V)
	}

	result := [][]string{}
	for _,v:= range ResultMap{
		result = append(result, v)
	}

	return result
}




func splitStringbyIndex(str string, i int,str2 string) (string) {
	// 优化，优先操作二进制数据
	rawStrSlice := []byte(str)
	sp1 := string(rawStrSlice[:i]) + str2
	sp2 := string(rawStrSlice[i:])
	// 直接操作字符串也可以
	//sp1 = str[:i]
	//sp2 = str[i:]
	return sp1+sp2
}