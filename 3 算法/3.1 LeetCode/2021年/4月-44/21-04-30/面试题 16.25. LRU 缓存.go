package main

import "fmt"

// 面试题 16.25. LRU 缓存
//设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存被填满时，它应该删除最近最少使用的项目。
//
//它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
//获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
//
//示例:
//
//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4

func main() {
	x := Constructor(2)
	x.Put(2,1)
	x.Get(2)
	x.Put(3,2)
}


type LRUCache struct {
	ThMap   map[int]*TowArr
	TheList *TowArr
	Cap     int
	Len     int
	TheLast *TowArr
}

type TowArr struct {
	Val  int
	Key  int
	Top  *TowArr
	Next *TowArr
}

func Constructor(capacity int) LRUCache {
	fmt.Println()
	x := LRUCache{
		ThMap:   map[int]*TowArr{},
		TheList: &TowArr{},
		Cap:     capacity,
		Len:     0,
	}
	x.TheLast = x.TheList
	return x
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.ThMap[key]; !ok {
		return -1
	}

	x := this.ThMap[key]
	top := x.Top
	next := x.Next

	if next!=nil{
		x.Top.Next = next
		next.Top = top
		this.TheLast.Next = x
		x.Top = this.TheLast
		this.TheLast = x
	}

	return this.ThMap[key].Val
}

func (this *LRUCache) Put(key int, value int) {
	// 值存在
	if _, ok := this.ThMap[key]; ok {
		this.ThMap[key].Val = value
		x := this.ThMap[key]
		this.TheLast.Next = x
		x.Top = this.TheLast
		this.TheLast = x
	}

	// 没有达到容量
	if this.Len < this.Cap {
		// 直接添加到末尾
		x := TowArr{
			Val: value,
			Key: key,
		}
		this.TheLast.Next = &x
		x.Top = this.TheLast
		this.TheLast = &x
		this.Len = this.Len + 1
		this.ThMap[key] = &x
	} else {
		TheOneVal := this.TheList.Next.Key
		delete(this.ThMap, TheOneVal)

		this.TheList = this.TheList.Next

		x := TowArr{
			Val: value,
			Key: key,
		}
		this.TheLast.Next = &x
		x.Top = this.TheLast
		this.TheLast = &x
		this.ThMap[key] = &x

	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
