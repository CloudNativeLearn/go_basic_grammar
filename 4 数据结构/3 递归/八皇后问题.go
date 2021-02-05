package main

import "fmt"

func main() {
	list := [8][7]int{}

	for i:=0;i<7;i++ {
		list[0][i] = 1
		list[7][i] = 1
	}
	for i:=0;i<8;i++ {
		list[i][0] = 1
		list[i][6] = 1
	}

	// 设置挡板1 表示
	list[3][1] = 1
	//list[3][2] = 1
	//list[1][2] = 1
	list[2][2] = 1
	SetWay(&list,1,1)
	for _,i :=range list{
		fmt.Println(i)
	}


}

func SetWay(list *[8][7]int , i int,j int) bool {
	if list[6][5]==2 {
		return true
	}else {
		if list[i][j] ==0{
			list[i][j] =2
			if SetWay(list,i+1,j){
				return true
			}else if SetWay(list,i,j+1) {
				return true
			}else if SetWay(list,i-1,j){
				return true
			}else if SetWay(list,i,j-1){
				return true
			}else {
				list[i][j] =3
				return false
			}
		}else {
			return false
		}
	}
}