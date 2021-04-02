package main
// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
//
//示例1：
//
//输入：阿里笔试->2->4, 阿里笔试->3->4
//输出：阿里笔试->阿里笔试->2->3->4->4
//限制：
//
//0 <= 链表长度 <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	
}

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 三个变量
	// l1 l1的头节点
	// l2 l2的头节点
	// NewNode  新链表的头节点
	if l2==nil {
		return l1
	}
	if l1==nil  {
		return l2
	}
	NewNode := &ListNode{}
	mid := NewNode
	for {


		if l1 == nil {
			mid.Next = l2
			break
		}
		if l2==nil {
			mid.Next = l1
			break
		}

		if  l1.Val<=l2.Val {
			mid.Next = l1
			mid=l1
			l1 = l1.Next
		}else {
			mid.Next = l2
			mid=l2
			l2 = l2.Next
		}



	}
	return NewNode.Next

}
