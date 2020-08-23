package main

import "fmt"

func sum(value ...int) int {
	total := 0
	for _, val := range value {
		total = total + val
	}

	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))

	values := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(sum(values...))
}
