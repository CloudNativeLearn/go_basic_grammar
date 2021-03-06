package main

import "sort"

// 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
//示例:
//
//输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
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
//链接：https://leetcode-cn.com/problems/group-anagrams
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func groupAnagrams(strs []string) [][]string {
	Thmap := map[string][]string{}
	for _, v := range strs {
		s := []byte(v)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sStr:= string(s)
		Thmap[sStr] = append(Thmap[sStr], v)
	}

	ans := make([][]string,len(Thmap))
	index :=0
	for _, v:= range Thmap{
		ans[index] = v
		index = index+1
	}
	return ans


}
