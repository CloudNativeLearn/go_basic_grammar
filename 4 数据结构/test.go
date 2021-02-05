package main

import "fmt"

func main() {
	a := &test{
		next: &test{
			next: &test{
				next: nil,
				id: 2,
			},
			id: 1,
		},
		id: 0,
	}
	for a!=nil{


		fmt.Printf("student=%+v\n", a)
		a = a.next


	}


}

type test struct {
	next *test
	id int
}
