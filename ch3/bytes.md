# ch3/bytes

#### bytes
1. 包含关系判断
    ```go
    // 判断字节切片是否包含子字节切片
    bytes.Contains([]byte("hello world"), []byte("hello"))  // true

    // 判断是否包含任意字节
    bytes.ContainsAny([]byte("hello"), []byte("hl"))  // true

    // 判断是否包含特定字节
    bytes.ContainsRune([]byte("hello"), 'h')  // true
    ```

2. 前缀/后缀判断
   ```go
    bytes.HasPrefix([]byte("hello world"), []byte("hello"))  // true
    bytes.HasSuffix([]byte("hello world"), []byte("world"))  // true
   ```

3. 子字节切片查找
    ```go
    // 查找子字节切片位置
    bytes.Index([]byte("hello"), []byte("l"))        // 2
    bytes.LastIndex([]byte("hello"), []byte("l"))    // 3

    // 查找字节位置
    bytes.IndexByte([]byte("hello"), 'e')    // 1
    bytes.IndexRune([]byte("hello"), 'l')    // 2
    ```

4. 大小写转换
   ```go
    bytes.ToUpper([]byte("Hello"))  // []byte("HELLO")
    bytes.ToLower([]byte("HELLO")) // []byte("hello")
    bytes.Title([]byte("hello world")) // []byte("Hello World")
   ```

5. 重复与替换
    ```go
    // 重复字节切片
    bytes.Repeat([]byte("go"), 3)  // []byte("gogogo")

    // 替换内容
    bytes.Replace([]byte("hello world"), []byte("world"), []byte("golang"), -1)
    bytes.ReplaceAll([]byte("hello world"), []byte("world"), []byte("golang"))
    ```

6. 分割字节切片
    ```go
    // 按空白分割
    bytes.Fields([]byte("hello world go"))  // [][]byte{[]byte("hello"), []byte("world"), []byte("go")}

    // 按特定分隔符分割
    bytes.Split([]byte("a,b,c"), []byte(","))      // [][]byte{[]byte("a"), []byte("b"), []byte("c")}
    bytes.SplitAfter([]byte("a,b,c"), []byte(",")) // [][]byte{[]byte("a,"), []byte("b,"), []byte("c")}
    ```

7. 连接字节切片
    ```go
    // 使用分隔符连接字节切片
    bytes.Join([][]byte{[]byte("a"), []byte("b"), []byte("c")}, []byte(","))  // []byte("a,b,c")
    ```

8. 去除空白/特定字节
    ```go
    // 去除首尾空白
    bytes.TrimSpace([]byte("  hello  "))  // []byte("hello")

    // 去除特定字节
    bytes.Trim([]byte("!!!hello!!!"), []byte("!"))  // []byte("hello")
    bytes.TrimLeft([]byte("!!!hello!!!"), []byte("!"))
    bytes.TrimRight([]byte("!!!hello!!!"), []byte("!"))
    ```

9. 计数比较
    ```go
    // 统计子字节切片出现次数
    bytes.Count([]byte("cheese"), []byte("e"))  // 3

    // 比较字节切片
    bytes.Compare([]byte("a"), []byte("b"))  // -1 (a < b)
    ```

10. 字节切片构建器
    ```go
    var builder bytes.Buffer
    builder.Write([]byte("Hello"))
    builder.Write([]byte(" "))
    builder.Write([]byte("World"))
    result := builder.Bytes()  // []byte("Hello World")
    ```

11. 字节切片与字符串转换
    ```go
    // 字符串转字节切片
    data := []byte("hello")

    // 字节切片转字符串
    str := string(data)

    // 高效的零拷贝转换
    str = unsafe.String(&data[0], len(data)) // unsafe但高效
    ```

12. 字节切片比较与操作
    ```go
    // 判断是否相等
    bytes.Equal([]byte("hello"), []byte("hello"))  // true

    // 判断前缀（高效版本）
    bytes.HasPrefix([]byte("hello world"), []byte("hello"))  // true

    // 读取直到分隔符
    reader := bytes.NewReader([]byte("hello\nworld"))
    line, _ := reader.ReadBytes('\n')  // []byte("hello\n")
    ```

