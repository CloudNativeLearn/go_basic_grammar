package main

func main() {

}

type LinkArr struct {
	val  int
	next *LinkArr
}

func GetByIndex(arr *LinkArr) *LinkArr {
	learn := 0
	head := arr
	for head.next != nil {
		learn = learn + 1
		head = head.next
	}
	number := (learn + 1) / 3
	head = arr
	midNumber := 0
	for {
		if midNumber == number {
			return head
		}
		midNumber = midNumber + 1
		head = head.next
	}
}
