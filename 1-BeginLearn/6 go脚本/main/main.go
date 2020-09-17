package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const all = "F:\\学习代码资料\\GoLang\\基础学习\\Go_Learning\\Go基础语法\\go_basic_grammar\\1-BeginLearn\\6 go脚本\\main\\all.txt"

const get = "F:\\学习代码资料\\GoLang\\基础学习\\Go_Learning\\Go基础语法\\go_basic_grammar\\1-BeginLearn\\6 go脚本\\main\\get.txt"

func getList(path string) *[]string {
	var StringList = []string{}
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件失败")
		return nil
	}



	defer f.Close()

	br := bufio.NewReader(f)
	for {
		s, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		StringList = append(StringList, string(s))
	}

	return &StringList
}

func main() {
	var all = getList(all)
	var get = getList(get)
	var s1 = []string{}

	for _, v1 := range *all {
		var not = ""
		for _, v2 := range *get {
			if v1 == v2 {
				not = v1
			}
		}
		if not == "" {
			s1 = append(s1, v1)
		}
	}

	fmt.Println(s1)

}
