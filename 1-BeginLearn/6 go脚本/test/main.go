package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
const all = "F:\\学习代码资料\\GoLang\\基础学习\\Go_Learning\\Go基础语法\\go_basic_grammar\\1-BeginLearn\\6 go脚本\\main\\all.txt"

func main() {

	var StringList = []string{}
	f, err := os.Open(all)
	if err != nil {
		fmt.Println("open error")
		return
	}
	defer f.Close()

	br := bufio.NewReader(f)
	for {
		s, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		StringList  = append(StringList,string(s))
	}
	fmt.Println(StringList)

}
