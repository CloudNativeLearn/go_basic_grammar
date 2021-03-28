package main

import "fmt"

func main() {
 list := [][]int{}
 a := []int{1,2,3}
 for i:=0;i<10;i++{
  list = append(list, make([]int,len(a)))
 }
 fmt.Println(list)
 s := "1,2,3"
 println(s[0])
 println(s[1])
}
