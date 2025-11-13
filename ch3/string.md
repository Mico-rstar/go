### go string特性
go中string本质上是一段**utf8编码**的**只读**的**字节序列**

#### unicode编码
unicode编码编码了世界上几乎所有符号，为每个符号都分配了一个32bit的码点。unicode码点在go中对应rune类型
unicode定长，所以用unicode存储字节序列可以进行O(1)的随机访问
但是缺点是浪费了很多空间

#### utf8编码
string使用utf8编码
utf8是变长的unicode表示
```go
0xxxxxxx                             runes 0-127    (ASCII)
110xxxxx 10xxxxxx                    128-2047       (values <128 unused)
1110xxxx 10xxxxxx 10xxxxxx           2048-65535     (values <2048 unused)
11110xxx 10xxxxxx 10xxxxxx 10xxxxxx  65536-0x10ffff (other values unused)
```
使用utf8编码string导致go中不能用索引i访问到第i个字符
```go
s := "中文"
fmt.Println(s[0])   // 不是 '中', 而是中的utf8编码的第一个字节
```
世界的等价表示
```go
"世界"  
"\xe4\xb8\x96\xe7\x95\x8c"  // 原始utf8字节序列
"\u4e16\u754c"  // unicode 16bit码点值
"\U00004e16\U0000754c"  // unicode 32bit码点值
```

#### 对字符进行操作
由于string只是字节序列，要对字符进行操作，需要把string转化为[]rune 类型
```go
// "program" in Japanese katakana
s := "プログラム"
fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
r := []rune(s)
fmt.Printf("%x\n", r)  // "[30d7 30ed 30b0 30e9 30e0]"
fmt.Println(string(r)) // "プログラム"
```


### []byte类型
string是定长的只读字节序列，而[]byte是定长的可写字节序列
#### string和[]byte的转换
```go
s := "abc"
b := []byte(s)  // 会拷贝一份数据
s2 := string(b) // 会拷贝一份数据
```



### 常用包
[[strings]]
[[bytes]]
[[strconv]]
[[unicode]]
    
### bytes.Buffer类型




