package main

import (
 "fmt"
 "reflect"
)

func main() {
 // 47 - 57
 // - 45
 // + 43
 // 空格 32
 //str := " -+01239"
 //for k,i:= range str{
 // fmt.Println(i)
 // fmt.Println(reflect.TypeOf(i))
 // fmt.Println(str[k])
 // //fmt.Println(string(i))
 //}

 s:="()"
 for k,v := range s{
  fmt.Println(k,v)
  fmt.Println(reflect.TypeOf(v))
 }
 a := []int32{1,2,3}
 fmt.Println(a[:len(a)-1])
}
