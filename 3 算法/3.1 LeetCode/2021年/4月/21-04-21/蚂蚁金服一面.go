package main
// //编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""。
//abc,abd,aqe123=>a
//efa,eqb,asdfasf=>""
import "fmt"

func main() {
	fmt.Println(GetString([]string{"a", "abc", "abd"}))
}

func GetString(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	st := arr[0]
	for i := 0; i < len(arr); i++ {
		st = GetCommon(st, arr[i])
	}
	return st
}

func GetCommon(a string, b string) string {
	prefinx := ""
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return prefinx
		}
		prefinx = prefinx + string(a[i])
	}
	return prefinx
}
