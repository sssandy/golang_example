package main

import(
	"fmt"
	"net/url"

	// "go.uber.org/zap"
)

func main()  {
	var urlStr string = "傻了吧:%:%@163& .html.html"
	escapeUrl := url.QueryEscape(urlStr)
	fmt.Println("编码:",escapeUrl)

	enEscapeUrl, _ := url.QueryUnescape(escapeUrl)
	fmt.Println("解码:",enEscapeUrl)
}