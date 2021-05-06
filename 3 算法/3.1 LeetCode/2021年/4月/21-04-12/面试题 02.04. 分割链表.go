package main

// 编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。
//
//示例:
//
//输入: head = 3->5->8->5->10->2->1, x = 5
//输出: 3->1->2->10->5->5->8
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/partition-list-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//func main() {
//	// 3->5->8->5->10->2->1
//	root := &ListNode{
//		Val: 3,
//		Next:&ListNode{
//			Val: 5,
//			Next: &ListNode{
//				Val: 8,
//				Next:&ListNode{
//					Val: 5,
//					Next: &ListNode{
//						Val: 10,
//						Next: &ListNode{
//							Val: 2,
//							Next: &ListNode{
//								Val: 1,
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//
//	partition(root,3)
//}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//3->5->8->5->10->2->1
func partition(head *ListNode, x int) *ListNode {
	xiao := &ListNode{}
	Da := &ListNode{}
	Xiao_in := &ListNode{}
	Da_in := &ListNode{}
	xiao.Next = Xiao_in
	Da.Next = Da_in
	for head != nil{
		if head.Val>=x{
			Da_in.Next = head
			Da_in = head
		}else{
			Xiao_in.Next = head
			Xiao_in = head
		}
		head  = head.Next
	}
	Xiao_in.Next = nil
	Da_in.Next = nil
	result:= AddTwo(xiao.Next.Next,Da.Next.Next)
	return result
}


// 合并两个链表
func AddTwo(One *ListNode,Two *ListNode) *ListNode {
	if One==nil{
		return Two
	}
	root := One
	for root.Next!=nil{
		root=root.Next
	}
	root.Next = Two
	return One


}