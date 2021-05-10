package main

// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
//
// 
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[4,5,1,2,3]
//示例 2：
//
//
//输入：head = [0,1,2], k = 4
//输出：[2,0,1]
// 
//
//提示：
//
//链表中节点的数目在范围 [0, 500] 内
//-100 <= Node.val <= 100
//0 <= k <= 2 * 109
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rotate-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func main() {
	a := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	rotateRight(&a, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	TheHead := head
	TheEnd := head
	TheMap := map[*ListNode]*ListNode{}
	for head.Next != nil {
		TheMap[head.Next] = head
		head = head.Next
	}
	TheMap[TheHead] = head
	TheEnd = head
	TheEnd.Next = TheHead
	for i := 0; i < k; i++ {
		TheHead = TheMap[TheHead]
		TheEnd = TheMap[TheEnd]
	}
	TheEnd.Next = nil
	return TheHead
}
