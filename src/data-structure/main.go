package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	// s := "hello"
	// sl := s[1:]

	// // 对切片进行修改
	// // 转化为[]byte
	// bs := []byte(s)
	// bs[1] = 'a'
	// fmt.Printf("sl = %v\n", sl)
	// fmt.Printf("bs = %v\n", string(bs))

	// // 直接定义为切片
	// a := []int{1, 2, 3, 4, 5}
	// reverse(a)
	// fmt.Printf("reverse a: %v\n", a)

	// // 利用切片来操作数组
	// arr := [...]int{1, 2, 3, 4}
	// reverse(arr[:])
	// fmt.Printf("利用切片反转后的arr: %v\n", arr)

	// str := "abcdefg"
	// sstr := str[1:2]
	// sstr = str[1:]
	// fmt.Printf("sstr = %v\n", sstr)

	s := make([]int, 5) // 创建一个cap=5, len=5的切片
	sl := s[:4]
	sl = append(sl, 1)
	fmt.Printf("s = %v\n", s)
	fmt.Printf("sl = %v\n", sl)


	// append
	// ss := "ab"
	// ssl := ss[:]
	// fmt.Printf("ssl = %v", ssl)
	// append(ssl, 'a', 'b')

	// ss := [...]int{1, 2}
	// ssl := ss[:]
	// ssl = append(ssl, 1)
	// fmt.Printf("ssl = %v", ssl)

	// ss := []int{1, 2}
	// ss = append(ss, 1)
	// fmt.Printf("ssl = %v", ss)

	// 对切片进行切片
	// ss := []int{1, 2, 3}
	// ssl := ss[1:]
	// ssl = append(ssl, 1)

	// ssl[0] = 100
	// fmt.Printf("ssl = %v\n", ssl)
	// fmt.Printf("ss = %v\n", ss)

	// 使用make
	// ssl := make([]int, 5)
	// ssl[2] = 200
	// fmt.Printf("len = %v\n", len(ssl))
	// fmt.Printf("cap = %v\n", cap(ssl))
	// sll := ssl[:2]
	// fmt.Printf("sll len = %v\n", len(sll))
	// fmt.Printf("sll cap = %v\n", cap(sll))
	// sll = append(sll, 1)
	// fmt.Printf("ssl = %v\n", ssl)
	// fmt.Printf("sll = %v\n", sll)


	// strs := []string{"monkey", " ", "malao"}
	// fmt.Printf("nonempty(strs) = %v\n", nonempty(strs))	
	// fmt.Printf("strs = %v\n", strs)


	test()
	testRefCopy()
}

func nonempty(strs []string) []string {
	out := strs[:0]
	for _, s := range strs {
		if s != " " {
			out = append(out, s)
		}
	}
	return out
}