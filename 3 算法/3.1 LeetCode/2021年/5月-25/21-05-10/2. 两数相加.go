package main

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 
//
//示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//示例 2：
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//示例 3：
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-two-numbers
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	TheResult := &ListNode{

	}
	TheEnd := TheResult.Next

	jinwei := 0
	val :=0
	for {
		if l1 == nil && l2 == nil {
			if jinwei!=0{
				TheEnd.Next=&ListNode{
					Val: jinwei,
				}
			}
			return TheResult.Next
		}

		if l1 != nil && l2 != nil {
			val = l1.Val + l2.Val + jinwei
			if val >= 10 {
				val = val % 10
				jinwei = 1
			}else {
				jinwei=0
			}
			if TheEnd == nil {
				TheResult.Next = &ListNode{
					Val: val,
				}
				TheEnd = TheResult.Next
			} else {
				TheEnd.Next = &ListNode{
					Val: val,
				}
				TheEnd = TheEnd.Next
			}
			l1 = l1.Next
			l2 = l2.Next
			continue
		}

		if l1 == nil {
			val = l2.Val + jinwei
			if val >= 10 {
				val = val % 10
				jinwei = 1
			}else {
				jinwei=0
			}
			if TheEnd == nil {
				TheResult.Next = &ListNode{
					Val: val,
				}
				TheEnd = TheResult.Next
			} else {
				TheEnd.Next = &ListNode{
					Val: val,
				}
				TheEnd = TheEnd.Next
			}
			l2 = l2.Next
		} else if l2 == nil {
			val = l1.Val + jinwei
			if val >= 10 {
				val = val % 10
				jinwei = 1
			}else {
				jinwei=0
			}
			if TheEnd == nil {
				TheResult.Next = &ListNode{
					Val: val,
				}
				TheEnd = TheResult.Next
			} else {
				TheEnd.Next = &ListNode{
					Val: val,
				}
				TheEnd = TheEnd.Next
			}
			l1 = l1.Next

		}

	}
}
