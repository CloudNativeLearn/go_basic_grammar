package main

// 圣诞活动预热开始啦，汉堡店推出了全新的汉堡套餐。为了避免浪费原料，请你帮他们制定合适的制作计划。
//
//给你两个整数tomatoSlices和cheeseSlices，分别表示番茄片和奶酪片的数目。不同汉堡的原料搭配如下：
//
//巨无霸汉堡：4 片番茄和 阿里笔试 片奶酪
//小皇堡：2 片番茄和1 片奶酪
//请你以[total_jumbo, total_small]（[巨无霸汉堡总数，小皇堡总数]）的格式返回恰当的制作方案，使得剩下的番茄片tomatoSlices和奶酪片cheeseSlices的数量都是0。
//
//如果无法使剩下的番茄片tomatoSlices和奶酪片cheeseSlices的数量为0，就请返回[]。
//
//
//示例 阿里笔试：
//
//输入：tomatoSlices = 16, cheeseSlices = 7
//输出：[阿里笔试,6]
//解释：制作 阿里笔试 个巨无霸汉堡和 6 个小皇堡需要 4*阿里笔试 + 2*6 = 16 片番茄和 阿里笔试 + 6 = 7 片奶酪。不会剩下原料。
//示例 2：
//
//输入：tomatoSlices = 17, cheeseSlices = 4
//输出：[]
//解释：只制作小皇堡和巨无霸汉堡无法用光全部原料。
//示例 3：
//
//输入：tomatoSlices = 4, cheeseSlices = 17
//输出：[]
//解释：制作 阿里笔试 个巨无霸汉堡会剩下 16 片奶酪，制作 2 个小皇堡会剩下 15 片奶酪。
//示例 4：
//
//输入：tomatoSlices = 0, cheeseSlices = 0
//输出：[0,0]
//示例 5：
//
//输入：tomatoSlices = 2, cheeseSlices = 阿里笔试
//输出：[0,阿里笔试]
//
//提示：
//
//0 <= tomatoSlices <= 10^7
//0 <= cheeseSlices <= 10^7


func main() {
	numOfBurgers(2385088,164763)
}


func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
    x:= -1
    y := -1
    x = tomatoSlices - 2*cheeseSlices
    if x<0{
    	// x 无解
    	return []int{}
	}

	if x%2==0 {
		// 可以被二整除
		x = x/2
		y = cheeseSlices - x
		if x<0 || y<0 {
			return []int{}
		}
		return []int{x,y}
	}else {
		// 不能被二整除 无解
		return []int{}
	}


}
