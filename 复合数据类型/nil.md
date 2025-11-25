### 什么类型的值可以为nil?
```go
// 1. 指针
var p *int = nil

// 2. 切片
var s []int = nil

// 3. 映射
var m map[string]int = nil

// 4. 通道
var ch chan int = nil

// 5. 函数
var f func() = nil

// 6. 接口: 当且仅当类型和值都为nil
var i interface{} = nil
```


### 理解
nil的本质是空指针，所以值可以为nil的类型，要么是引用类型，要么底层是指针来实现