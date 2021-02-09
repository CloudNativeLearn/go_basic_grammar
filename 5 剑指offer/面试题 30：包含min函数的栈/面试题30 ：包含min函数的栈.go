package main

func main() {

}

type MinStack struct {
	TheDataStack []int
	MinDataStack []int
	MinData      int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		TheDataStack: []int{},
		MinDataStack: []int{},
		MinData: -1,
	}
}

func (this *MinStack) Push(x int) {
// 判断长度是否为空
	if len(this.TheDataStack)==0 {
		// 如果为空  给minDataStack 和 MinData赋值
		this.MinData = x
		this.TheDataStack = append(this.TheDataStack,x)
		this.MinDataStack = append(this.MinDataStack,x)
	}else {
		if x<this.MinData {
			this.MinData = x
			this.MinDataStack = append(this.MinDataStack,x)
			this.TheDataStack = append(this.TheDataStack,x)
		}else {
			this.MinDataStack = append(this.MinDataStack,this.MinData)
			this.TheDataStack = append(this.TheDataStack,x)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.TheDataStack)!=0 {

		this.MinDataStack = this.MinDataStack[:len(this.MinDataStack)-1]
		this.TheDataStack = this.TheDataStack[:len(this.TheDataStack)-1]
		if len(this.MinDataStack)!=0 {
			this.MinData = this.MinDataStack[len(this.MinDataStack)-1]
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.TheDataStack)!=0 {
		return this.TheDataStack[len(this.TheDataStack)-1]
	}else {
		return -1
	}
}

func (this *MinStack) Min() int {
return this.MinData
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
