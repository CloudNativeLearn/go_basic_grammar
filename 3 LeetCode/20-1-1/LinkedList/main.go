package main

import "fmt"

// 给你一个链表的头节点 head，请你编写代码，反复删去链表中由 总和 值为 0 的连续节点组成的序列，直到不存在这样的序列为止。
//
//删除完毕后，请你返回最终结果链表的头节点。
//
// 
//
//你可以返回任何满足题目要求的答案。
//
//（注意，下面示例中的所有序列，都是对 ListNode 对象序列化的表示。）

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}


func removeZeroSumSublists(head *ListNode) *ListNode {
	sentinel := &ListNode{0,head} //哨兵节点
	sumMap := map[int]*ListNode{}
	cur := sentinel
    //  阿里笔试 2 -2 2
	// 初次遍历存储sum和节点的对应关系
	//  阿里笔试  0  2  0 3
	sum := 0
	for cur!=nil{
		sum += cur.Val
		sumMap[sum] = cur
		cur = cur.Next
	}

	fmt.Println(sentinel)

	// 第二次遍历节点，进行节点删除操作
	cur = sentinel
	sum = 0
	for cur!=nil{
		sum += cur.Val
		if sumMap[sum]!= cur{
			cur.Next = sumMap[sum].Next
		}
		cur = cur.Next
	}


	return sentinel.Next
}

func main() {
	head := &ListNode{
		1,
		nil,
	}

	head.Next = &ListNode{
		2,
		nil,
	}
	head.Next.Next=&ListNode{
		-2,
		nil,
	}
	head.Next.Next.Next=&ListNode{
		2,
		nil,
	}
	removeZeroSumSublists(head)

}


//func removeZeroSumSublists(head *ListNode) *ListNode {
//	 theMap := map[int]*ListNode{}
//
//	 //第一次遍历 链表转hash表
//	 thesum :=0
//	 for p := head;p!=nil;p=p.Next{
//	 	thesum = p.Val + thesum
//	 	theMap[thesum] = p
//	 }
//
//	 // 第二次遍历 删除节点
//	for p:= head;p!=nil;p=p.Next {
//		thesum = p.Val + thesum
//		p.Next = theMap[thesum].Next
//	}
//
//
//
//}
