package main

import "fmt"

func main() {
	x := [][]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12}}
	fmt.Println(XUnhuan(&x,len(x[0]),len(x)))
	
}
// [[阿里笔试,2,3,4],[5,6,7,8],[9,10,11,12]]

// list 表示数组   xlen表示x长度  ylen表示y长度
func XUnhuan( list *[][]int,xLen int,yLen int)[]int {
	if xLen ==0 || yLen==0{
		return []int{}
	}
	start := 0
	Alllist := []int{}
	for xLen>start*2 && yLen>start*2{
		ShowInfo(list,xLen,yLen,start,&Alllist)
		start = start + 1
	}
	return Alllist
}

// 输出函数
// list 总数组
// xlen x的长度
// ylen y的长度
// start 开始长度
// Alllist 输出数组
func ShowInfo(list *[][]int,xLen int,yLen int,start int,Alllist *[]int)  {
x := xLen-start-1
y := yLen-start-1
// 从左向右
	for i:=start;i<=x;i++ {
		*Alllist=append(*Alllist, (*list)[start][i])
	}
// 从右向下
	if start<y {
		for i:=start+1;i<=y;i++ {
			*Alllist = append(*Alllist,(*list)[i][x])
		}
	}

// 从右向左
	if  start<y && start<x{
		for i:=x-1;i>=start;i-- {
			*Alllist = append(*Alllist,(*list)[y][i])
		}
	}

// 从左到上
	if start<x && start<y-1 {
		for i:=y-1;i>=start+1;i-- {
			*Alllist = append(*Alllist,(*list)[i][start])
		}
	}
}




