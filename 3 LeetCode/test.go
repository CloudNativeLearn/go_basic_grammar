package main

import "fmt"

func main() {

	xxx:= "((()))"
	bbb:= ""
	for k,V := range xxx{
		fmt.Println(k)
		fmt.Println(V)
	}
	fmt.Println(xxx[0]=='(')
	fmt.Println(len(xxx))
	fmt.Println(xxx[3:4]+bbb)
	fmt.Println(xxx[0:len(
		xxx)-1])


}
