package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// 斐波那契数列 1, 2, 3 ,5
func Fibonacci(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("请输入正数")
	}
	if n == 1 || n == 2 {
		return n, nil
	}
	a, _ := Fibonacci(n - 1)
	b, _ := Fibonacci(n - 2)
	return a + b, nil
}

// 实现一个简单的拓扑排序
// 使用Kahn算法
func TopoSort(m map[string][]string) ([]string, error) {
	order := make([]string, 0, 10)
	// 创建一个map记录每个节点的入度
	indu := make(map[string]int)
	for k, vl := range m {
		indu[k] = len(vl)
		for _, node := range vl {
			indu[node]++
		}
	}
	// 将所有入度为0的node加入队列
	zeroQueue := make([]string, 0, 10)
	for k, v := range indu {
		if v == 0 {
			zeroQueue = append(zeroQueue, k)
		}
	}
	for len(zeroQueue) != 0 {
		node := zeroQueue[0]
		zeroQueue = zeroQueue[1:]
		order = append(order, node)
		for _, nd := range m[node] {
			indu[nd]--
			if indu[nd] == 0 {
				zeroQueue = append(zeroQueue, nd)
			}
			if indu[nd] < 0 {
				return nil, errors.New("unexpected behaviour")
			}
		}
	}
	if len(order) != len(indu) {
		return nil, errors.New("不存在合法拓扑排序序列")
	}
	return order, nil

}

func ArgTest() {
	s := []int{1, 2, 3}
	// 计算切片平均值优雅写法
	avg := func(nums []int) int {
		sum := 0
		for _, i := range nums {
			sum += i
		}
		return sum / len(nums)
	}(s)
	fmt.Printf("avg = %d\n", avg)
}

func TopoSortTest() {
	// prereqs记录了每个课程的前置课程
	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}
	order, err := TopoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to toposort: %v\n", err)
	}
	for _, token := range order {
		fmt.Printf("%v, ", token)
	}
	fmt.Println()
}




func MutableLenArgsTest() {
	a := 1
	b := "ss"
	fmt.Printf("%d, %s\n", a, b)

	accumulate := func (vals ...int) int {
		sum := 0
		for _, v := range vals {
			sum += v
		}
		return sum
	}
	fmt.Println(accumulate(1, 2, 3))
}



// 闭包值捕获行为
func ClosureTest() {
	// a := 1
	// // 作为参数传入闭包，值传递
	// f1 := func (num int) {
	// 	num = 2
	// }
	// f1(a)

	// // 直接在环境中捕获，引用语义
	// f2 := func ()  {
	// 	a = 2
	// }
	// f2()
	// fmt.Println(a)

    var funcs []func()
    
    for i := 0; i < 3; i++ {
        // 将匿名函数添加到切片中，该函数捕获了循环变量 i
        funcs = append(funcs, func() {
            fmt.Println(i) // 注意：这里捕获的是变量 i 本身
        })
    }
    
    // 循环结束后，再依次执行这些函数
    for _, f := range funcs {
        f()
    }
}


// 使用defer写一个函数用于统计函数运行时间
func timming() func(){
	start := time.Now()
	return func () {
		fmt.Printf("执行时长：%#v\n", time.Since(start))
	}
}

func testTimming() {
	defer timming()()
	for i := 0; i < 100; i++ {
		fmt.Printf("i = %d, ", i)
	}
	fmt.Println()
}

func DeferPrint() {
	defer func ()  {
		fmt.Println("defer的函数被调用")
	}()
	fmt.Println("正常函数调用")
}


func DeferTest() {
	testTimming()
	DeferPrint()
}