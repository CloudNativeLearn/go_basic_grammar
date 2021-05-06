package main

import "sort"

// 假设你正在读取一串整数。每隔一段时间，你希望能找出数字 x 的秩(小于或等于 x 的值的个数)。请实现数据结构和算法来支持这些操作，也就是说：
//
//实现 track(int x) 方法，每读入一个数字都会调用该方法；
//
//实现 getRankOfNumber(int x) 方法，返回小于或等于 x 的值的个数。
//
//注意：本题相对原题稍作改动
//
//示例:
//
//输入:
//["StreamRank", "getRankOfNumber", "track", "getRankOfNumber"]
//[[], [1], [0], [0]]
//输出:
//[null,0,null,1]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rank-from-stream-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


type StreamRank struct {
	nums []int
}

func Constructor() StreamRank {
	return StreamRank{}
}

func (sr *StreamRank) Track(x int) {
	i := sort.Search(len(sr.nums), func(i int) bool {
		return sr.nums[i] >= x
	})
	sr.nums = append(append(sr.nums[:i:i], x), sr.nums[i:]...)
}

func (sr *StreamRank) GetRankOfNumber(x int) int {
	return sort.Search(len(sr.nums), func(i int) bool {
		return sr.nums[i] > x
	})
}


