package main


// 给定两个用链表表示的整数，每个节点包含一个数位。
//
//这些数位是反向存放的，也就是个位排在链表首部。
//
//编写函数对这两个整数求和，并用链表形式返回结果。
//
// 
//
//示例：
//
//输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
//输出：2 -> 1 -> 9，即912
//进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?
//
//示例：
//
//输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
//输出：9 -> 1 -> 2，即912
//通过次数26,198提交次数56,356
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-lists-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
		},


	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 9,
		},


	}
	addTwoNumbers(l1,l2)
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	Beizhu := 1

	Jinwei := 0
	root := &ListNode{}
	root_in := &ListNode{}
	root.Next = root_in
	WhetherInwei := false

	for l1 != nil || l2 != nil {
		l1Value := 0
		if l1 != nil {
			l1Value = l1.Val
			l1 = l1.Next
		}
		l2Value := 0
		if l2 != nil {
			l2Value = l2.Val
			l2 = l2.Next
		}
		MidResult := l1Value + l2Value + Jinwei
		if MidResult >= 10 {
			gewei := MidResult % 10
			//Result = Result + gewei*Beizhu
			Jinwei = MidResult / 10
			WhetherInwei = true
			root_in.Next = &ListNode{
				Val: gewei,
			}
			root_in = root_in.Next
		} else {
			//Result = Result + MidResult*Beizhu
			Jinwei = 0
			WhetherInwei = false
			root_in.Next = &ListNode{
				Val: MidResult,
			}
			root_in = root_in.Next
		}
		Beizhu = Beizhu * 10
	}
	//Result = Result + Beizhu*Jinwei
	//root_in.Next = &ListNode{
	//	Val: Beizhu*Jinwei,
	//}
	if WhetherInwei==true {
		root_in.Next = &ListNode{
			Val: Jinwei,
		}
	}



	return root.Next.Next

}
