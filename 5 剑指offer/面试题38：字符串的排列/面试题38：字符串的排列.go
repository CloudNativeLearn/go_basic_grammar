// 请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。
//
//输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
//输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

package main

import "fmt"

func main() {
	permutation("abc")
}


func permutation(s string) []string {

	StringList := []byte(s)
	fmt.Println(StringList)
	result := []string{}
	digui(0, &StringList, &result)
	fmt.Println(result)
	return result
}

func digui(y int, StringList *[]byte, result *[]string) {
	if y == len(*StringList)-1 {
		x := *StringList
		*result = append(*result, string(x[:]))
		fmt.Println(11)
		return
	}

	TheMap := map[string]int{}
	for i:=y;i<len(*StringList);i++ {
		if _, ok := TheMap[string((*StringList)[i])]; ok {
			continue
		}
		TheMap[string((*StringList)[i])]=0
		mid := (*StringList)[i]
		(*StringList)[i] = (*StringList)[y]
		(*StringList)[y] = mid


		digui(y+1, StringList, result)

		mid = (*StringList)[y]
		(*StringList)[y] = (*StringList)[i]
		(*StringList)[i] = mid


	}

}



