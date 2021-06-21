package main

// 对链表进行插入排序。
//
//
//插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
//每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。
//
// 
//
//插入排序算法：
//
//插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
//每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
//重复直到所有输入数据插入完为止。
// 
//
//示例 1：
//
//输入: 4->2->1->3
//输出: 1->2->3->4
//示例 2：
//
//输入: -1->5->3->4->0
//输出: -1->0->3->4->5
//通过次数85,590提交次数126,877
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/insertion-sort-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func main() {
	insertionSortList(&ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	},)
}
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	TheHead := head
	TheMind := head
	TheMind = nil
	Head := &ListNode{
		Next: head,
	}
	for TheHead.Next!=nil{
		if TheHead.Next.Val<TheHead.Val{
			TheMind = TheHead.Next
			TheHead.Next = TheHead.Next.Next
			// 从head遍历到TheHead 进行插入
			// 判断第一个
			if Head.Next.Val>TheMind.Val{
				TheMind.Next = Head.Next
				Head.Next = TheMind

			}else {
				y := Head.Next
				for y.Next!=nil{
					if y.Next.Val>TheMind.Val{
						TheMind.Next = y.Next
						y.Next = TheMind
						break
					}
					y=y.Next

				}
			}

		}else {
			TheHead = TheHead.Next
		}


	}



	return Head.Next
}
