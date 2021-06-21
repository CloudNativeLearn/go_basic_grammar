package main

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//进阶：你能尝试使用一趟扫描实现吗？
//
// 
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//示例 2：
//
//输入：head = [1], n = 1
//输出：[]
//示例 3：
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

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

func main() {
	list := &ListNode{
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
	removeNthFromEnd(list,2)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	theHead := &ListNode{}
	theHead.Next = head
	theHead2 := theHead
	s := Stark{list: []*ListNode{}}
	for theHead2 != nil {
		s.push(theHead2)
		theHead2 = theHead2.Next
	}
	var a *ListNode
	for i := 0; i <= n; i++ {
		a = s.pop()
	}
	a.Next = a.Next.Next
	return  theHead.Next

}

type Stark struct {
	list []*ListNode
}

func (this *Stark) pop() *ListNode {
	a := this.list[len(this.list)-1]
	this.list = this.list[:len(this.list)-1]
	return a
}

func (this *Stark) push(a *ListNode) {
	this.list = append(this.list, a)
}
