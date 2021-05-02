package main

// 面试题 17.08. 马戏团人塔
//有个马戏团正在设计叠罗汉的表演节目，一个人要站在另一人的肩膀上。出于实际和美观的考虑，在上面的人要比下面的人矮一点且轻一点。已知马戏团每个人的身高和体重，请编写代码计算叠罗汉最多能叠几个人。
//
//示例：
//
//输入：height = [65,70,56,75,60,68]
//    weight = [100,150,90,190,95,110]
//输出：6
//解释：从上往下数，叠罗汉最多能叠 6 层：(56,90), (60,95), (65,100), (68,110), (70,150), (75,190)
//提示：
//
//height.length == weight.length <= 10000

func main() {
	bestSeqAtIndex([]int{65,70,56,75,60,68},[]int{100,150,90,190,95,110})
}


func bestSeqAtIndex(height []int, weight []int) int {
	if len(height)==0{
		return 0
	}
	QuickSortStark(0, len(height)-1, height, weight)
	ThResult := 0
	result := []int{}
	for i := 0; i < len(weight); i++ {
		result = append(result, 1)
		for j := 0; j < i; j++ {
			if weight[i] > weight[j] && height[i] > height[j] {
				result[i] = GetMax(result[i],result[j]+1)
				ThResult = GetMax(result[i],ThResult)
			}
		}
	}
	return ThResult
}

func GetMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetTheMaxInARR(arr []int)int  {
	Thmax := arr[0]
	for _,v := range arr{
		if v>Thmax{
			Thmax = v
		}
	}
	return Thmax
}
func QuickSortStark(l int, r int, arr []int, arr2 []int) {
	x := (l + r) / 2
	Theleft := l
	Theright := r
	Xvalue := arr[x]
	for l < r {

		// 在右边找到比 xvalue小的值
		for arr[l] < Xvalue {
			l = l + 1
		}

		// 在左边找到比 xvalue大的值
		for arr[r] > Xvalue {
			r = r - 1
		}

		if l >= r {
			break
		}

		temp := arr[r]
		arr[r] = arr[l]
		arr[l] = temp

		temp = arr2[r]
		arr2[r] = arr2[l]
		arr2[l] = temp

		if arr[l] == Xvalue {
			r = r - 1
		}
		if arr[r] == Xvalue {
			l = l + 1
		}

	}

	if l == r {
		l = l + 1
		r = r - 1
	}

	if Theleft < l {
		QuickSortStark(Theleft, r, arr, arr2)
	}
	if Theright > r {
		QuickSortStark(l, Theright, arr, arr2)
	}
}
