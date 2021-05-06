package main

import "fmt"

type HeroNode struct {
	Number   int       // 排名
	Name     string    // 名字
	NickName string    // 别名
	NextNode *HeroNode // 下一个节点
}

type SingleLinkedList struct {
	Head *HeroNode
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		Head: &HeroNode{
			Number:   0,
			Name:     "",
			NickName: "",
			NextNode: nil,
		},
	}
}

// 在最后添加节点
func (this *SingleLinkedList) AddHeroNode(Her *HeroNode) {
	mid := this.Head

	if mid.NextNode == nil {
		mid.NextNode = Her
		return
	}
	for mid.NextNode != nil {
		mid = mid.NextNode
	}
	// mid 为最后一个
	mid.NextNode = Her
}

// 在指定位置添加节点
func (this *SingleLinkedList) AddByindex(Her *HeroNode, Index int) {
	mid := this.Head
	if mid.NextNode==nil{
		mid.NextNode = Her
	}else {
		for mid.NextNode!=nil{
				if mid.NextNode.Number==Index{
					fmt.Println("该排名已存在")
					return
				}else if mid.NextNode.NextNode==nil {
					fmt.Println(11)
					mid.NextNode.NextNode = Her
					return
				}
				mid = mid.NextNode
		}
	}


}

// 删除指定节点
func (this *SingleLinkedList) DeleteByIndex(Index int) {
	mid := this.Head
	for mid != nil {
		if mid.NextNode == nil {
			fmt.Println("到达链表最后，指定元素不存在")
			return
		}
		if mid.NextNode.Number == Index {
			//x:= mid.NextNode
			mid.NextNode = mid.NextNode.NextNode
		}
		mid = mid.NextNode
	}
}

//  显示所有节点
func (this *SingleLinkedList) ShowAllNode() {
	mid := this.Head
	for mid != nil {
		fmt.Printf("HeroNode=%+v\n", mid)
		mid = mid.NextNode
	}
}

// 链表 迭代反转链表
func (this *SingleLinkedList) IterativeInversion() {
	midx := this.Head
	if midx.NextNode == nil{
		fmt.Println("节点个数为0 无法反转")
	}else {
		var beg *HeroNode
		mid := midx.NextNode
		end :=midx.NextNode.NextNode
		head := mid
		for  {

			if end==nil {
				break
			}
			mid.NextNode = beg
			beg = mid
			mid = end
			end=end.NextNode
		}
		mid.NextNode = beg
		head = mid
		for head != nil {
			fmt.Printf("HeroNode=%+v\n", head)
			head = head.NextNode
		}
	}
}

// 头插法反转链表
func (this *SingleLinkedList)ReverseLinked()  {
	head := this.Head
	if head.NextNode==nil {
		fmt.Println("节点个数为0 ")
	}
	NewHead := &HeroNode{}
	for  {
		if head == nil{
			break
		}
		if NewHead.NextNode == nil{
			x := head.NextNode
			NewHead.NextNode = head
			head.NextNode = nil
			head = x
		}else {
			x:= head.NextNode
			y := NewHead.NextNode
			NewHead.NextNode = head
			head.NextNode = y
			head=x
		}


		fmt.Println(99)
	}

	for NewHead != nil {
		fmt.Printf("HeroNode=%+v\n", NewHead)
		NewHead = NewHead.NextNode
	}
}

// 4、就地逆置法反转链表
func (this *SingleLinkedList)LocalReverse()  {
	head := this.Head
	if head.NextNode == nil || head.NextNode.NextNode==nil {
		fmt.Println("节点个数为0或者1")
	}
	beg := head.NextNode
	end := head.NextNode.NextNode
	for  {

		x := head.NextNode
		y := end.NextNode

		head.NextNode = end
		end.NextNode = x

		beg.NextNode = y
		end = y
		if beg.NextNode==nil {
			break
		}

	}
	for head != nil {
		fmt.Printf("HeroNode=%+v\n", head)
		head = head.NextNode
	}
}

