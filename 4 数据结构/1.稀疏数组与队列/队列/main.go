package main

import "fmt"

func main() {
	queue := CreateArrayQueue(10)
	fmt.Println(queue)
	SetMenu(queue)
	
}

func SetMenu(arr *ArrayQueue)  {

	for {
		fmt.Println("=========队列选择菜单==========")
		fmt.Println("1 向队列添加数据")
		fmt.Println("2 取出队列数据")
		fmt.Println("3 显示队列所有数据")
		fmt.Println("4 输出队列第一个数据")
		fmt.Println("5 退出")
		fmt.Println("选择对应的数字")
		var input int
		fmt.Scanln(&input)
		fmt.Println(input)
		switch input {
		case 1:
			fmt.Println("输入数据")
			fmt.Scanln(&input)
			arr.AddValue(input)
			fmt.Println("添加成功")
		case 2:
			fmt.Printf("输出的数据为 %d \n",arr.GetValue())
		case 3:
			arr.ShowAllArr()
		case 4:
			fmt.Printf("第一个数据为 %d \n",arr.ShowTheFirst())
		case 5:
			fmt.Println("在建")
			return
		}

	}
}
