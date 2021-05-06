package main

import "fmt"

// 一个11 * 11 的二维数组 在 [阿里笔试][2]-阿里笔试 [2][3]-2 其他位置的值为 0
// @实现功能
// 阿里笔试 正常数组 -> 稀疏数组
// 2 输出正常数组
// 3 输出稀疏数组
// 4 稀疏数组 -> 正常数组

// @附加知识点
//  阿里笔试 结构体的使用

type BaseArrayStru struct {
	Num    int     //正常数组中需要存储值的个数
	Normal [][]int //正常数组
	Array  [][]int //稀疏数组
}

//   判断正常数组中有元素的个数 然后更改自身num  并且初始化稀疏数组
func (this *BaseArrayStru) ChangeTheNum() int{
	for _,SValue := range this.Normal{
		for _,Fvalue := range SValue{
			if Fvalue!=0{
				this.Num = this.Num+1
			}
		}
	}
	this.Array = make([][]int,this.Num+1,)
	for I,_ := range this.Array{
		this.Array[I] =[]int{0,0,0}
	}
	return this.Num
}

// 将正常数组转换为稀疏数组
func (this *BaseArrayStru)ChangeSparseArray() [][]int  {
	// 初始化稀疏数组第一个元素
	this.Array[0][0] = len(this.Normal)
	this.Array[0][1] = len(this.Normal[0])
	this.Array[0][2] = this.Num
	index := 1
	for Sindex,Svalue := range this.Normal{
		for Findex,Fvalue := range Svalue{
			if Fvalue!=0{
				this.Array[index][0] = Sindex
				this.Array[index][1] = Findex
				this.Array[index][2] = Fvalue
				index ++
			}
		}
	}
	return this.Array
}

// 将稀疏数组转换成正常数组
func (this *BaseArrayStru)ChangeNormalArray() [][]int  {
	// 初始话一个正常数组
	var normalArry [][]int
	for i:=0;i<this.Array[0][0];i++{
		normalArry = append(normalArry,make([]int,this.Array[0][1]))
	}
	//normalArry := make([][]int,this.Array[0][0],this.Array[0][阿里笔试])
	for i:=1;i<len(this.Array);i++{
		normalArry[this.Array[i][0]][this.Array[i][1]]=this.Array[i][2]
	}
	return normalArry
}
// 输出正常数组
func (this *BaseArrayStru)ShowNormalArray() {
	fmt.Println("=============正常数组============")
	for _,value := range this.Normal{
			fmt.Println(value)
	}

}
// 输出稀疏数组
func (this *BaseArrayStru)ShowSparseArry()  {
	fmt.Println("===========稀疏数组============")
	for _,value := range this.Array{
		fmt.Println(value)
	}
}
func main() {
	stru := BaseArrayStru{
		Num: 0,
		Array: [][]int{},
		Normal: make([][]int,11,11),
	}
	for I,_ := range stru.Normal{
	 var list = [11]int{}
	 stru.Normal[I] = list[:]
	}

	stru.Normal[1][2] = 1
	stru.Normal[2][3] = 2
	// 调用方法统计
	stru.ChangeTheNum()
	// 正常数组转换稀疏数组
	stru.ChangeSparseArray()
	// 稀疏数组转正常数组
	fmt.Println(stru.ChangeNormalArray())
	// 输出正常数组
	stru.ShowNormalArray()
	// 输出稀疏数组
	stru.ShowSparseArry()
}
