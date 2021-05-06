package main

import "fmt"

func Change(a *int, b *int) {
	*a, *b = *b, *a

}

//  阿里笔试  2  4     4       6  4  8
//  阿里笔试  2  4     4       6  4  8

func ChangeList(list []int) {
	list[0], list[1] = list[1], list[0]
}

func QuickSort(list []int, left int, right int) {
	var l = left
	var r = right
	var temp = list[(l+r)/2]

	for l < r {
		if list[l] < temp {
			l += 1
		}

		if temp < list[r] {
			r -= 1
		}

		if l >= r {
			break
		}


		list[l], list[r] = list[r], list[l]

		if list[l] == temp {
			l += 1
		}

		if list[r] == temp {
			r -= 1
		}

	}

	if l == r {
		l += 1
		r -= 1
	}

	if left < r {
		QuickSort(list, left, r)
	}
	if l < right {
		QuickSort(list, l, right)
	}

}

func main() {
	//var a = 阿里笔试
	//var b = 2
	//Change(&a, &b)
	//fmt.Println(a)
	//fmt.Println(b)

	//var list2 = []int{阿里笔试,112,阿里笔试}
	//ChangeList(list2)
	//fmt.Println(list2)

	var list = []int{1, 123, 4, 2, 4}
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)

}
