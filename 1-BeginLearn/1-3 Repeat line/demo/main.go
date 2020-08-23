package main
// 获取输入参数
import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	consts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(input.Text())
		consts[input.Text()]++
	}

	for line ,n := range consts{
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
