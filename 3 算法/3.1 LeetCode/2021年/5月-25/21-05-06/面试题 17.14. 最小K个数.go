package main

// 设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。
//
//示例：
//
//输入： arr = [1,3,5,7,2,4,6,8], k = 4
//输出： [1,2,3,4]
//提示：
//
//0 <= len(arr) <= 100000
//0 <= k <= min(100000, len(arr))
//通过次数29,365提交次数54,046
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/smallest-k-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


func smallestK(arr []int, k int) []int {
QuickSortStark(0,len(arr)-1,arr)
return arr[:k]
}


func QuickSortStark(l int,r int,arr []int)  {
	x := (l+r)/2
	Theleft := l
	Theright := r
	Xvalue := arr[x]
	for l<r{

		// 在右边找到比 xvalue小的值
		for arr[l]<Xvalue{
			l = l+1
		}

		// 在左边找到比 xvalue大的值
		for  arr[r]>Xvalue {
			r = r-1
		}

		if l>=r {
			break
		}

		temp := arr[r]
		arr[r] = arr[l]
		arr[l] = temp

		if arr[l] == Xvalue {
			r = r-1
		}
		if arr[r] == Xvalue{
			l = l+1
		}

	}

	if l ==r {
		l = l+1
		r = r-1
	}

	if Theleft<l {
		QuickSortStark(Theleft,r,arr)
	}
	if Theright>r {
		QuickSortStark(l,Theright,arr)
	}
}
