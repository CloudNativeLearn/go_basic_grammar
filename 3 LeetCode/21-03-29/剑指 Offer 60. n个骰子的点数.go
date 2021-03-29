package main

import "fmt"

// 把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
//
// 
//
//你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。
//
// 
//
//示例 1:
//
//输入: 1
//输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]
//示例 2:
//
//输入: 2
//输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(dicesProbability(3))
}
//
//func dicesProbability(n int) []float64 {
//	Result := map[int]float64{}
//	x := map[int]float64{}
//	if n ==0 {
//		return []float64{}
//	}else {
//		Result[1],Result[2] ,Result[3],Result[4],Result[5],Result[6] = 1.0/6,1.0/6,1.0/6,1.0/6,1.0/6,1.0/6
//		x = Stark(n-1,&Result)
//	}
//	for k,v :=range x{
//		fmt.Println(k,v)
//	}
//return nil
//}
//
//// 递归处理
//func Stark(n int,Map *map[int]float64) map[int]float64 {
//	NewMap := map[int]float64{}
//	if n!=0 {
//		for i:=1;i<=6;i++{
//			for k,v := range *Map{
//				//fmt.Println(k,v)
//				NewMap[k+i] = NewMap[k+i]+ v*1.0/6
//			}
//		}
//		Map = &NewMap
//		n = n-1
//		if n!=0 {
//			Stark(n,Map)
//		}
//	}
//	return NewMap
//}

func dicesProbability(n int) []float64 {

	x := []float64{1.0/6,1.0/6,1.0/6,1.0/6,1.0/6,1.0/6}
	if n ==0{
		return []float64{}
	}else {
		Stark(n-1,&x)
	}
	return x
}
func Stark(n int,arr *[]float64) {
NewArr := []float64{}

if n!=0{
	for i:=0;i<6;i++{
		for j,v := range *arr{
			if i+j<len(NewArr) {
				//NewArr = append(NewArr, (NewArr)[i+j] + v*1.0/6)
				(NewArr)[i+j] = (NewArr)[i+j]+v*1.0/6
				//*arr = append(*arr, 0)
			}else {
				NewArr = append(NewArr, v*1.0/6)
				//*arr = append(*arr, 0)
			}
		}
	}
	fmt.Println(NewArr)
	*arr = make([]float64,len(NewArr))
	copy(*arr,NewArr)
	Stark(n-1,arr)
}
}


