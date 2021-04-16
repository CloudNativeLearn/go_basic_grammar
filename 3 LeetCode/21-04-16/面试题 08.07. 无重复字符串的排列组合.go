package main

// 无重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合，字符串每个字符均不相同。
//
//示例1:
//
// 输入：S = "qwe"
// 输出：["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
//示例2:
//
// 输入：S = "ab"
// 输出：["ab", "ba"]
//提示:
//
//字符都是英文字母。
//字符串长度在[1, 9]之间。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/permutation-i-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//func main() {
//	fmt.Println(permutation("qwe"))
//}
//

func permutation(S string) []string {
	b := []byte(S)
	result := &[]string{}
	MindString := ""
	Doit(b,result,MindString)
	return *result
}

func Doit(arr []byte,result *[]string,MindString string)  {
	if len(arr)==0 {
		*result = append(*result, MindString)
	}

	for K,V := range  arr{
		NewString := MindString + string(V)
		top := make([]byte,len(arr[:K]))
		copy(top,arr[:K])
		Newarr := append(top,arr[K+1:]...)
		Doit(Newarr,result,NewString)
	}
}
