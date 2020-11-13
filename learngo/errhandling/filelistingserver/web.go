package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/captainlee1024/learngo/errhandling/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

// errWrapper 处理 handler 里返回的错误信息，handler 只进行逻辑的处理，遇到错误就返回出来，由这里统一接收处理
// 函数式编程 输入是函数，输出也是函数,把输入的函数包装一下输出
func errWrapper(
	handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		// panic
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Panic: %v\n", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		// user error
		if err != nil {
			fmt.Printf("Error handling request: %s\n",
				err.Error())

			if userError, ok := err.(userError); ok {
				http.Error(writer,
					userError.Message(),
					http.StatusBadRequest)
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound // TDOD: 找不到文件返回状态码 404
			case os.IsPermission(err):
				code = http.StatusForbidden // TODO: 没有权限返回状态码 403
			default:
				code = http.StatusInternalServerError // TODO: 未知错误返回转状态码 500
			}
			http.Error(writer,
				http.StatusText(code),
				code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	// http.HandleFunc("/list/",
	// 	errWrapper(filelisting.HandleFileList))

	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
