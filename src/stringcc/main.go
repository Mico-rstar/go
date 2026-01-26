package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "asd"
	s2 := "bcs"
	s3 := "lll"
	// 创建新变量，拷贝原字符串，O(n^2)
	c1 := s1 + s2 + s3

	// 由于使用接口参数（泛型），需要在运行时判断数据类型，要用到反射机制，效率比+还要低
	c2 := fmt.Sprintf("%s%s%s", s1, s2, s3)

	// 使用 []byte 缓冲区，可以提前分配空间，也可以依赖自动扩容，O(n)
	var b strings.Builder
	b.Grow(100)
	b.WriteString(s1)
	b.WriteString(s2)
	b.WriteString(s3)
	c3 := b.String()

	// 由strings.Builder实现，性能相近
	c4 := strings.Join([]string{s1, s2, s3}, "")

	fmt.Printf("+: %s\nSprintf: %s\nstrings.Builder: %s\nstrings.Join: %s\n", c1, c2, c3, c4)

}
