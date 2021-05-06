package main

import "fmt"

func main() {
	SliceDelete()
}

func SliceDelete() {
	// 初始化一个新的切片 seq
	seq := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(seq[1:])

	// 指定删除位置
	index := 6

	// 输出删除位置之前和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	// seq[index+阿里笔试:]... 表示将后段的整个添加到前段中
	// 将删除前后的元素连接起来
	seq = append(seq[:len(seq)-1], seq[len(seq):]...)
	// 输出链接后的切片
	fmt.Println(seq)
}