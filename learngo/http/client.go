package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	// resp, err := http.Get("http://blog.leecoding.club")
	// resp, err := http.Get("https://captainlee1024.gitee.io/blog")

	request, err := http.NewRequest(http.MethodGet,
		"http://www.imooc.com", nil)
	// request.Header.Add("User-Agent",
	// 	"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	// resp, err := http.DefaultClient.Do(request)

	client := http.Client{
		CheckRedirect: func(req *http.Request,
			via []*http.Request) error {
			fmt.Println("Redirect : ", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	// 两个返回值 resp *http.Response, err error
	// 源码注释提示：Caller should close resp.Body when done reading from it.需要关闭它
	// resp, err := http.Get("http://www.imooc.com")
	// resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}
