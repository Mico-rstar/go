### new函数
new函数在go中和变量声明语句的内存分配方式，也就是说只是变量声明一个语法糖
```Go
// 直接声明
var p int
// 使用new
p := new(int)
```
```Go
type Person struct {
    Name    string
    Age     int
}
psn := new(Person)
```
```Go
// new返回指针
p := new(int)
fmt.Printf("%T\n", p)
```

#### 与java和cpp中new的区别
java和cpp中的new都是堆内存分配原语，是一个关键字
go中的new只是一个预定义函数，并不能决定内存最终是分配在堆上还是栈上，取决于编译器的逃逸分析

### make函数
make函数专门用于创建切片、映射、通道三种引用类型
```Go
// 创建切片, len = 5, cap = 10
s := make([]int, 5, 10)

// 创建map
m := make(map[string]int)

// 创建缓冲大小为5的通道
c := make(chan int, 5)
```

### new 和 make对比
#### 共通点
- new和make都用于创建引用类型


#### 不同点
- new可以创建所有类型，但new不适合初始化引用类型，因为new会直接将其进行0值初始化
  ```Go
  // new 只做了这一步：
    sliceHeader := &reflect.SliceHeader{
        Data: 0,  // 底层数组指针为 nil
        Len:  0,  
        Cap:  0,
    }
  ```

- make专门用来初始化slice, map, channel, make会进行更复杂的初始化，包括为底层数据结构申请空间，初始化len和cap等
  




