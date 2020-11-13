package main

import (
	"fmt"
	// "io"
	"time"

	"github.com/captainlee1024/learngo/retriever/mock"
	"github.com/captainlee1024/learngo/retriever/real"
)

// Retriever download 需要的接口，该接口有 download 需要的 Get 方法
type Retriever interface {
	Get(url string) string
}

// Poster 拥有 Post 方法的接口
type Poster interface {
	Post(url string,
		form map[string]string) string
}

const url = "http://www.baidu.com"

func download(r Retriever) string {
	// return r.Get("http://www.baidu.com")
	// return r.Get("http://blog.leecoding.club")
	// return r.Get("http://www.imooc.com")
	// return r.Get("https://captainlee1024.gitee.io/blog")
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
			map[string]string{
				"name":"user1",
				"todo":"test",
			})
}

// RetrieverPoster 组合接口，组合 Retriever 的 Get 方法和 Poster 的 Post 方法
type RetrieverPoster interface{
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url,
		map[string]string{
			"contents":"another faked baidu.com",
		})
	return s.Get(url)
}
// inspect 查看接口信息
func inspect(r Retriever) {
	fmt.Println("inspecting : ", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch :")
	switch v := r.(type) {
	case *mock.MyRetriever:
		fmt.Println("Contents :", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent :", v.UserAgent)
	}
	fmt.Println()
}

func main()  {
	// 接口里面有两块儿东西
	// 1.类型
	// 2.值
	var r Retriever
	mockRetriever := mock.MyRetriever{Contents: "this is a fake baidu.com"}
	// 因为参数是值，所以这里是拷贝整个 mock.MyRetriever 到 r

	// 指针传递,放地址进入
	r = &mockRetriever
	fmt.Println(r)
	inspect(r)
	// fmt.Printf("%T %v\n", r, r)
	// fmt.Print(download(r))
	// fmt.Println(download(
	// 	mock.MyRetriever{
	// 		Context: "this is a fake baudu.com"},
	// ))

	// 参数是指针，所以这里拷贝 real.Retriever 的地址和值的地址
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	// fmt.Printf("%T %v\n\n", r, r)
	// fmt.Println(download(r))

	// 查看接口信息
	// 1.
	inspect(r)

	// 2.type assertion
	if mockRetriever, ok := r.(*mock.MyRetriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.UserAgent, realRetriever.TimeOut)
	} else {
		fmt.Println("not a real retriever")
	}
	fmt.Println("Try a session")
	fmt.Print(session(&mockRetriever))
}