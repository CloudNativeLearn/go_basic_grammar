package main

import "fmt"

var TheParent = map[int]int{}
func main() {
EEE("")
}

func EEE(s string)  {
	if s =="2" {
		return
	}else {
		for k,v := range "121"{
			fmt.Println(k,string(v))
			EEE(string(v))
		}
	}
}
