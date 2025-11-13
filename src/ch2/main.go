package main

import (
	"fmt"
)

func main() {
	s := "str好啦"
	n := len(s)
	for i := 0; i < n; i++ {
		fmt.Println(s[i])
	}

	// for i := 0; i < n; {
	// 	r, size := utf8.DecodeRuneInString(s[i:])
	// 	fmt.Printf("%d\t%c\n", i, r)
	// 	i += size
	// }

	for i, r := range s {
		fmt.Printf("%d\t%c\n", i, r)
	}

	rs := "麻烦了"
	r := []rune(rs)
	fmt.Println(r)
	fmt.Println(string(r))

}
