package main

import "fmt"

func main() {
	arr := []int{1, 4, 6, 2, 1}
	ChooseSort(arr)
}

//func getLeastNumbers(arr []int, k int) []int {
//
//}

// 选择排序
func ChooseSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		temp := arr[min]
		arr[min] = arr[i]
		arr[i] = temp
	}
	fmt.Println(arr)
}

// 插入排序
func InsertSort(arr []int) {

}
