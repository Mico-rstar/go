package main

import (
	"fmt"
)

func PanicTest(i int) {
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		panic("unexpected arguments")
	}
}

// defer修改返回结果
func Num2String(n int) (result string) {
	defer func ()  {
		result = fmt.Sprintf("( %s )", result)
	}()
	return fmt.Sprintf("%d", n)
}

func EnhanceTest() {
	PanicTest(1)
	PanicTest(2)
	// defer printStack()
	// PanicTest(3)
	RecoverFromPanic()
	
	fmt.Println("修改后的字符串：", Num2String(2))
}




// recover from panic
func RecoverFromPanic() {
	defer func ()  {
		if p := recover(); p !=nil {
			fmt.Println("Recovered from panic:", p)
		}
	}()
	// panic("unexpected")
	PanicTest(3)
	fmt.Println("not be print")
}