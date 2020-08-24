package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type (
	// gResult  映射到从搜索拿到的结果文档
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}
	// gResponse包含定级的文档
	gRespinse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		}  `json:"responseData"`
	}
)

func main() {


	// 读取文件的内容
	content, err := ioutil.ReadFile("./json.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))


	// 将json响应解码到结构类型
	var gr gRespinse
	//err = json.NewDecoder(content).Decode(&gr)
	json.Unmarshal([]byte(string(content)), &gr)

	fmt.Println(gr)

}
