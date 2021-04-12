package main

import "fmt"

func main() {

a :=12
fmt.Println(a/10)
fmt.Println(a%10)

}

type test struct {
	next *test
	id int
}
