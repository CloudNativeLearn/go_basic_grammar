package main

import (
	"fmt"
)

func main() {
	rc := RongYunUtil()
	userRongyun, err := rc.UserRegister("1379773664584536064", "fyl", "http://img3.downza.cn/download/202008/152201-5f33989911385.jpg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userRongyun.Token)
}
