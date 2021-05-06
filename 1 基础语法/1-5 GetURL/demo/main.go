package main
// 获取URL

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _,url := range os.Args[1:]{

			// 链接url 检测是否出错
			resp,err := http.Get(url)
			if err!= nil{
				fmt.Fprintf(os.Stderr,"fetch: %v\\n",err)
				os.Exit(1)
			}

			// 读取网页的内容
			b,err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err!= nil{
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
				os.Exit(1)
			}
			fmt.Printf("%s",b)



	}
}
