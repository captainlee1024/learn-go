package main

import (
	"fmt"
	"regexp"
)

const text = `My email is abcde123@qq.com.acb safdas@qq.com.cn 
email1 is 644052732@qq.com
email2 is  xdfadfs@gmail.org
email3  is asdfa1@163.com
`

func main() {

	// 会返回两个值，正则表达式匹配器，如果不符合正则表达式返回想用的错误
	// re, err := regexp.Compile("644052732@qq.com")

	// 只返回正则表达式的匹配器
	// 一般需要我们在程序中直接写一些东西的时候可以使用该函数，因为一般写代码我们认为我们写的肯定不会错的
	// re := regexp.MustCompile("abcde123@qq.com")
	// re := regexp.MustCompile(".+@.+\\..+")
	// re := regexp.MustCompile(`.+@.+\..+`)
	// re := regexp.MustCompile(`[a-zA-Z0-9]+@.+\..+`)
	// re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)

	// re 拿到之后我们就可以在字符串中去找匹配上面的正则表达式的东西
	// match := re.FindString(text)

	// match := re.FindAllString(text, -1)

	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
	// fmt.Println(match)
}
