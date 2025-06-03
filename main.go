package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 待匹配的字符串
	str := "9:/opt/"

	// 编译正则表达式：匹配数字+冒号+任意内容
	re := regexp.MustCompile(`^(\d+):(.*)$`)
	// 查找匹配项
	match := re.FindStringSubmatch(str)
	if len(match) > 1 {
		// match[0] 是完整匹配 "9:"
		// match[1] 是捕获的第一个分组 "9"
		fmt.Println("提取的数字:", match[1])
		fmt.Println("提取的数字hou:", match[2])
	} else {
		fmt.Println("未找到匹配项")
	}
}
