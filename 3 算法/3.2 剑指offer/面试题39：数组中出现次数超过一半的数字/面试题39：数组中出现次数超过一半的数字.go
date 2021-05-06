package main

func main() {

}
func majorityElement(nums []int) int {
  TheMap := map[int]int{}
	for _,v :=range nums{
		if _,ok := TheMap[v];ok {
			TheMap[v] = TheMap[v] +1
		}else {
			TheMap[v] = 1
		}
	}

	var TheNumber int
	ThValue:=0
	for k,v := range TheMap{
		if TheNumber == 0{
			TheNumber = v
			ThValue = k
		}

		if TheNumber<v {
			TheNumber = v
			ThValue = k
		}else {
			ThValue=ThValue
		}
	}
return ThValue
}