package main

import "fmt"

func main() {
	list := []int{1, 3, 2, 6, 5}
	fmt.Println(verifyPostorder(list))

}

func verifyPostorder(postorder []int) bool {
	return Whether(postorder, len(postorder))
}

func Whether(postorder []int, length int) bool {
	if length == 0 || length < 0 {
		return true
	}

	// 头节点
	root := postorder[length-1]

	// 找到那个比i大的位置
	i := 0
	for ; i < length-1; i++ {
		if postorder[i] > root {
			break
		}

	}

	// 判断右边 是不是全部比root大
	j := i
	for ; j < length-1; j++ {
		if postorder[j] < root {
			return false
		}
	}

	// 判断左边是否为二叉搜索树
	left := true
	if i > 0 {
		left = Whether(postorder, i)
	}
	// 判断右边是否为二叉搜索树
	right := true

	if i < length-1 {
		right = Whether(postorder[i:length-1], length-i-1)
	}
	return right && left

}
