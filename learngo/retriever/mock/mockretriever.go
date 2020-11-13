package mock

import "fmt"

// MyRetriever 测试实现接口
type MyRetriever struct{
	Contents string
}

// Get 实现 Get 方法
func (r *MyRetriever) Get(url string) string {
	return r.Contents
}

// 作为接口实现者，我们可以实现多个接口的方法

// Post 实现 Post 方法（实现接口 Poster 的方法）
func (r *MyRetriever) Post(url string,
	form map[string]string) string {
		r.Contents = form["contents"]
	return "ok"
}

// Stringer 实现 fmt 包 的Stringer() 方法
func (r *MyRetriever) String() string {
	return fmt.Sprintf(
		"MyRetriever: {Contents = %s}", r.Contents)
}