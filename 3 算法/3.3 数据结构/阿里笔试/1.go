package main

import "fmt"

func main() {
	Function(3,4,[]string{"n","k","6","9"})
}

func Function(m int, n int, arr []string) {
	DataMap := map[string]int{}
	//// 排序arr
	//for i := 0; i < len(arr); i++ {
	//	for j := i; j < len(arr); j++ {
	//		if arr[j] > arr[i] {
	//			temp := arr[j]
	//			arr[j] = arr[i]
	//			arr[i] = temp
	//		}
	//	}
	//}
	fmt.Println(arr)
	for _, v := range arr {
		DataMap[v] = DataMap[v] + 1
	}
	fmt.Println(DataMap)
	// m>n 返回null
	if m > n {
		fmt.Println("不存在")
	}

	//利用递归
	str := ""
	newresuld := []string{}
	FunStart(DataMap,str,&newresuld,m)
	fmt.Println(newresuld)


}

// 递归函数 递归组合密码
func FunStart(TheMap map[string]int, str string, Result *[]string, m int) {
	// 如果字符串合法
	if  PassWhetherTrue(str,m){
		Result = AddList(Result,str)
	}else {
		// 如果字典不为空 循环map继续迭代继续递归
		if  MapWhetherEmpty(TheMap){
			for k,v := range TheMap{
				if v!=0 {
					str = str + k
					NewMap := CopyMap(TheMap)
					NewMap[k] = NewMap[k]-1
					FunStart(NewMap,str,Result,m)
				}

			}
		}
	}


}


// 添加数据不存在不添加
func AddList(arr *[]string,str string)  *[]string{
	had := false
	for _, v := range *arr{
		if v==str{
			had = true
		}
	}
	if !had {
		*arr = append(*arr, str)
	}
	fmt.Println("===")
	fmt.Println(*arr)
	return arr
}
// 判断map是否为空
func MapWhetherEmpty(tmap map[string]int) bool {
	w := true
	for _, v := range tmap {
		if v != 0 {
			w = true
		}
	}
	return w
}

// 判断密码串是否合法
func PassWhetherTrue(str string, m int) bool {
	chaNumber := 0
	NumNumber := 0
	for _,k := range str {
		if k >= 47 && k <= 57 {
			NumNumber = NumNumber + 1
		} else {
			chaNumber = chaNumber + 1
		}
	}
	return chaNumber >= 2 && NumNumber >= 1 && m == len(str)
}

// 拷贝数组
func CopyMap(oldMap map[string]int) map[string]int {
	NewMap := map[string]int{}
	for k, v := range oldMap {
		NewMap[k] = v
	}
	return NewMap
}
