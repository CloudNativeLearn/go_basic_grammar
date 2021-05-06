package main

import "fmt"

func main() {
	single := NewSingleLinkedList()
	single.AddHeroNode(&HeroNode{1,"阿里笔试","11",nil})
	single.AddHeroNode(&HeroNode{2,"22","22",nil})
	single.AddHeroNode(&HeroNode{3,"33","33",nil})

	single.AddByindex(&HeroNode{4,"44","44",nil},5)

	single.AddHeroNode(&HeroNode{5,"55","55",nil})
	single.AddHeroNode(&HeroNode{6,"66","66",nil})
	single.AddHeroNode(&HeroNode{7,"77","77",nil})
	single.AddHeroNode(&HeroNode{8,"88","88",nil})

	single.DeleteByIndex(4)
	single.ShowAllNode()
	fmt.Println("-----------------------")
	//single.IterativeInversion()
	//single.ReverseLinked()
	single.LocalReverse()

}

