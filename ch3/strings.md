# ch3/strings

#### strings
1. 包含关系判断
    ```go
    // 判断字符串是否包含子串
    strings.Contains("hello world", "hello")  // true

    // 判断是否包含任意字符
    strings.ContainsAny("hello", "hl")  // true

    // 判断是否包含特定字符
    strings.ContainsRune("hello", 'h')  // true
    ```
2. 前缀/后缀判断
   ```go
    strings.HasPrefix("hello world", "hello")  // true
    strings.HasSuffix("hello world", "world")  // true
   ```
   
3. 子串查找
    ```go
    // 查找子串位置
    strings.Index("hello", "l")        // 2
    strings.LastIndex("hello", "l")    // 3

    // 查找字符位置
    strings.IndexByte("hello", 'e')    // 1
    strings.IndexRune("hello", 'l')    // 2
    ```

4. 大小写转换
   ```go
    codestrings.ToUpper("Hello")  // "HELLO"
    strings.ToLower("HELLO") // "hello"
    strings.Title("hello world") // "Hello World"
   ```
   
5. 重复与替换
    ```go
    // 重复字符串
    strings.Repeat("go", 3)  // "gogogo"

    // 替换内容
    strings.Replace("hello world", "world", "golang", -1)
    strings.ReplaceAll("hello world", "world", "golang")
    ```
    
6. 分割字符串
    ```go
    // 按空格分割
    strings.Fields("hello world go")  // ["hello", "world", "go"]

    // 按特定分隔符分割
    strings.Split("a,b,c", ",")      // ["a", "b", "c"]
    strings.SplitAfter("a,b,c", ",") // ["a,", "b,", "c"]
    ```
7. 连接字符串
    ```go
    // 使用分隔符连接字符串切片
    strings.Join([]string{"a", "b", "c"}, ",")  // "a,b,c"
    ```
    
8. 去除空白/特定字符
    ```go
    // 去除首尾空白
    strings.TrimSpace("  hello  ")  // "hello"

    // 去除特定字符
    strings.Trim("!!!hello!!!", "!")  // "hello"
    strings.TrimLeft("!!!hello!!!", "!")
    strings.TrimRight("!!!hello!!!", "!")
    ```
9. 计数比较
    ```go
    // 统计子串出现次数
    strings.Count("cheese", "e")  // 3

    // 比较字符串
    strings.Compare("a", "b")  // -1 (a < b)
    ```
10. 字符串构建器
    ```go
    var builder strings.Builder
    builder.WriteString("Hello")
    builder.WriteString(" ")
    builder.WriteString("World")
    result := builder.String()  // "Hello World"
    ```
