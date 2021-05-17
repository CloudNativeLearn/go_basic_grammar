package main

import (
	"fmt"
)

//"52","3","v","71","J","A","0","v",       4    4
//"51","E","k","H","96","21","W","59",     4    4
// "I","V","s","59","w","X","33","29",     3    5
// "H","32","51","f","i","58","56","66",   5    3
//"90","F","10","93","53","85","28","78",  7    1
//"d","67","81","T","K"                    2     3
func main() {
	a := []int{0,1,2,3,4,5}
	fmt.Println(a[1:2])
}
//
//func deleteDuplicates(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	dummy := &ListNode{0, head}
//
//	cur := dummy
//	for cur.Next != nil && cur.Next.Next != nil {
//		if cur.Next.Val == cur.Next.Next.Val {
//			x := cur.Next.Val
//			for cur.Next != nil && cur.Next.Val == x {
//				cur.Next = cur.Next.Next
//			}
//		} else {
//			cur = cur.Next
//		}
//	}
//
//	return dummy.Next
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return nil
	}

	TheHead := &ListNode{}
	TheHead.Next = head
	head2 := TheHead
	for head2.Next!= nil && head2.Next.Next !=nil{
		if head2.Next.Val == head2.Next.Next.Val{
			x := head2.Next.Val
			for head2.Next!=nil && head2.Next.Val==x{
				head2.Next = head2.Next.Next
			}
		}else {
			head2 = head2.Next
		}
	}
	return TheHead.Next
}
