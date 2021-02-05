package main

import "fmt"

type ArrayQueue struct {
	maxSize int   //表示最大
	front   int   //队列头
	rear    int   //队列尾
	arr     []int // 用于存放数据
}

// 1 初始化队列    ！！不绑定在对象上
func CreateArrayQueue(max int) *ArrayQueue {
	return &ArrayQueue{
		maxSize: max,
		front:   0,
		rear:    0,
		arr:     make([]int, max),
	}
}

// 2 判断队列是否已满
func (this *ArrayQueue) WhetherFull() bool {
	//return this.front == this.maxSize-1
	return (this.rear+1)%this.maxSize==this.front;
}

// 3 判断队列是否为空
func (this *ArrayQueue) isEmpty() bool {
	return this.front == this.rear
}

// 4 添加数据到队列
func (this *ArrayQueue) AddValue(value int) {
	if !this.WhetherFull() {
		//this.rear = this.rear + 1
		this.arr[this.rear] = value
		this.rear = (this.rear+1)%this.maxSize
	}
	fmt.Errorf("队列已满，无法添加")

}

// 5 取出数据到队列
func (this *ArrayQueue) GetValue() int {
	if this.isEmpty() {
		fmt.Println("队列为空")
		fmt.Errorf("抱歉 队列为空")
		return -1
	}
	this.front = this.front + 1
	
	return this.arr[this.front]
}

// 6 显示队列所有数据
func (this *ArrayQueue) ShowAllArr() {
	fmt.Println("===========显示所有的数据===========")
	//for i := range this.arr {
	//	fmt.Println(i)
	//}
	fmt.Println(this.arr)
}

// 7 显示队列的头数据
func (this *ArrayQueue) ShowTheFirst() int {
	return this.arr[this.front+1]
}
