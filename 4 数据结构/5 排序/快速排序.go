package main

import "fmt"

func main() {
arr := []int{1,5,1,8,10,22,1}
QuickSortStark(0,len(arr)-1,arr)
fmt.Println(arr)
}


func QuickSortStark(l int,r int,arr []int)  {
  x := (l+r)/2
  Theleft := l
  Theright := r
  Xvalue := arr[x]
	for l<r{

		// 在右边找到比 xvalue小的值
		for arr[l]<Xvalue{
			l = l+1
		}

		// 在左边找到比 xvalue大的值
		for  arr[r]>Xvalue {
			r = r-1
		}

		if l>=r {
		break
		}

		temp := arr[r]
		arr[r] = arr[l]
		arr[l] = temp

		if arr[l] == Xvalue {
			r = r-1
		}
		if arr[r] == Xvalue{
			l = l+1
		}

	}

	if l ==r {
	l = l+1
	r = r-1
	}

	if Theleft<l {
		QuickSortStark(Theleft,r,arr)
	}
	if Theright>r {
		QuickSortStark(l,Theright,arr)
	}
}
