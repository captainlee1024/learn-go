package main

import "fmt"

// 最长不含有重复字符的子串长度
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start + 1 > maxLength {
			maxLength = i-start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	m := map[string]string {
		"name": "小王",
		"site": "www.xx.com",
		"quality": "notbad",
	}
	fmt.Println(m)

	m2 := make(map[string]int) //m2 == empty map
    fmt.Println(m2)
    
    var m3 map[string]int //m3 == nil
	fmt.Println(m3)
	
	for k, v := range m {
		fmt.Println(k, v)
	}

	name2 := m["nam"] //输入错误的 key
	fmt.Println(name2)

	name := m["name"]
	fmt.Println(name)

	if name, ok := m["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key does not exist")
	}

	delete(m, "name")
	name, ok := m["name"] //上面定义过了，这里要用 = 
	fmt.Println(name, ok)

	fmt.Println(
		lengthOfNonRepeatingSubStr("aaaabbc"),
		lengthOfNonRepeatingSubStr("abcabcd"),
		lengthOfNonRepeatingSubStr("a"),
		lengthOfNonRepeatingSubStr(""),
	)
}