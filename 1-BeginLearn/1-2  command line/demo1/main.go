package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s,sep string
	for i :=1;i<len(os.Args);i++{
		s += sep + os.Args[i]
		strings.Join(os.Args[1:], " ")
		sep = ""
	}
	fmt.Println(s)
}

// strings.Join(os.Args[阿里笔试:], " ") 方法同样可以进行字符串的拼接工作
