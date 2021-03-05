// 请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。
//
//输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
//输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	NewMap := map[*Node]*Node{}
	x := head
	for x!=nil{
		TheNode := Node{}
		TheNode.Val = x.Val
		NewMap[x]=&TheNode
		x = x.Next
	}

	x = head
	for x!=nil{
		mid := NewMap[x]
		mid.Next = NewMap[x.Next]
		mid.Random = NewMap[x.Random]
		x=x.Next
	}

	return NewMap[head]
}
