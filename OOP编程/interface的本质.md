# 方法/interface的本质

### interface的底层数据结构
#### eface 
**eface是空接口interface{}的底层数据结构**，只存储原类型信息和底层数据引用信息
![alt text](image.png)
```go 
type eface struct {
    _type   *_type
    data    unsafe.Pointer
}
```


#### iface
**iface是带方法interface的底层数据结构**，包含itab和data两部分
![alt text](image-1.png)
```go
type iface struct {
    tab     *itab
    data    unsafe.Pointer
}

type itab struct {
    inter   *interfacetype
    _type   *_type
    hash    uint32
    _       [4]byte
    fun     [x]uintptr     // 接口有几个方法x就是几
}
```


### 赋值行为
#### 赋值值类型
赋值值类型会创建对象副本，可能触发堆分配，由逃逸分析确定
```go
d := Dog{Name: "Buddy"}
s = d
d.Name = "Max"
fmt.Println(s.Speak()) // 输出: Woof! I'm Buddy （副本未变）
```

#### 赋值指针类型
赋值指针类型，只是赋值指针，增加了对底层数据的引用
```go
d := &Dog{Name: "Buddy"}
s = d
d.Name = "Max"
fmt.Println(s.Speak()) // 输出: Woof! I'm Max （共享同一对象）
```



### 通过interface机制如何实现动态分发？
1. 将不同类型变量赋值给接口类型变量
    ```go
    // 将具体类型放入接口类型切片中
    al := []Animal{
    Eagle{"老鹰"},
    Snake{"蛇"},
    }
    ```
    
2. 通过接口类型调用接口方法
    ```go
    for _, animal := range al {
		Battle(king, animal)
	}
    ```
    

3. 运行时动态分发
    运行时会通过 iface -> itab -> func 的链路动态分发方法
    ![alt text](image-1.png)