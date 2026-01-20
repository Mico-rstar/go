# ch3/bytes.Buffer

#### bytes.Buffer
1. 创建和初始化
    ```go
    // 创建空的Buffer
    var buf1 bytes.Buffer
    buf2 := new(bytes.Buffer)
    buf3 := bytes.NewBuffer([]byte("initial"))

    // 从字符串创建Buffer
    buf4 := bytes.NewBufferString("hello")
    ```

2. 写入数据
    ```go
    var buf bytes.Buffer

    // 写入字节切片
    buf.Write([]byte("hello "))     // 返回写入的字节数和错误

    // 写入单个字节
    buf.WriteByte('W')              // 返回错误

    // 写入字符串
    buf.WriteString("world")        // 返回写入的字节数和错误

    // 写入Rune
    buf.WriteRune('中')             // 返回写入的字节数和错误

    // 从Reader读取并写入
    buf.ReadFrom(strings.NewReader(" data")) // 返回读取的字节数和错误
    ```

3. 读取数据
    ```go
    var buf bytes.Buffer
    buf.WriteString("hello world")

    // 读取到字节切片
    data := make([]byte, 5)
    n, err := buf.Read(data)       // data = []byte("hello"), n = 5

    // 读取单个字节
    b, err := buf.ReadByte()        // b = ' '

    // 读取单个Rune
    r, size, err := buf.ReadRune()  // r = 'w', size = 1

    // 读取直到分隔符
    line, err := buf.ReadBytes('d') // 返回包含分隔符的字节切片
    text, err := buf.ReadString('d') // 返回包含分隔符的字符串
    ```

4. 获取数据
    ```go
    var buf bytes.Buffer
    buf.WriteString("hello")

    // 获取字节切片
    data := buf.Bytes()             // []byte("hello")

    // 获取字符串
    str := buf.String()             // "hello"

    // 注意：返回的数据不是副本，修改会影响Buffer内容
    data[0] = 'H'                  // buf的内容也会被修改为 "Hello"
    ```

5. 缓冲区管理
    ```go
    var buf bytes.Buffer

    // 获取已用字节数
    buf.WriteString("hello")
    length := buf.Len()             // 5

    // 获取总容量
    capacity := buf.Cap()           // >= 5

    // 确保最小容量
    buf.Grow(100)                   // 确保容量至少为100+当前长度

    // 重置缓冲区
    buf.Reset()                     // 清空内容，保留容量
    ```

6. 截断和扩展
    ```go
    var buf bytes.Buffer
    buf.WriteString("hello world")

    // 截断到指定长度
    buf.Truncate(5)                 // 保留前5个字节 "hello"

    // 扩展缓冲区（填充零值）
    buf.Truncate(10)                // 扩展到10字节，后面用0填充
    ```

7. 下一个字节操作
    ```go
    var buf bytes.Buffer
    buf.WriteString("hello")

    // 查看下一个字节但不移除
    b, err := buf.Next(1)           // b = []byte("h")

    // 查看多个字节但不移除
    data := buf.Next(3)             // data = []byte("ell")

    // 剩余内容仍可在后续读取中获取
    ```

8. 高效操作
    ```go
    // 预分配容量避免频繁扩容
    buf := bytes.NewBuffer(make([]byte, 0, 1024))

    // 重用Buffer
    func processData() {
        var buf bytes.Buffer
        buf.Grow(512) // 预分配

        // 写入数据
        buf.WriteString("result")

        // 使用结果
        result := buf.String()

        // Buffer会自动重置，不需要手动操作
        _ = result
    }
    ```

9. 与IO接口集成
    ```go
    // 作为io.Writer
    func writeToBuffer(w io.Writer) {
        w.Write([]byte("data"))
    }

    var buf bytes.Buffer
    writeToBuffer(&buf)             // Buffer实现了io.Writer接口

    // 作为io.Reader
    func readFromBuffer(r io.Reader) {
        data := make([]byte, 10)
        r.Read(data)
    }

    buf.WriteString("hello world")
    readFromBuffer(&buf)            // Buffer实现了io.Reader接口
    ```

10. 实用示例
    ```go
    // 字符串构建
    func buildString(parts []string) string {
        var buf bytes.Buffer
        buf.Grow(len(strings.Join(parts, ""))) // 预分配

        for _, part := range parts {
            buf.WriteString(part)
        }
        return buf.String()
    }

    // JSON编码（简化版）
    func encodeJSON(data map[string]interface{}) ([]byte, error) {
        var buf bytes.Buffer
        buf.WriteByte('{')

        first := true
        for key, value := range data {
            if !first {
                buf.WriteByte(',')
            }
            first = false

            buf.WriteByte('"')
            buf.WriteString(key)
            buf.WriteByte('"')
            buf.WriteByte(':')

            // 简化的值处理
            buf.WriteString(fmt.Sprintf("%v", value))
        }

        buf.WriteByte('}')
        return buf.Bytes(), nil
    }

    // 数据压缩（概念演示）
    func compressData(input []byte) []byte {
        var buf bytes.Buffer
        buf.Grow(len(input) / 2) // 假设压缩率为50%

        // 简化的RLE压缩
        i := 0
        for i < len(input) {
            count := 1
            for i+count < len(input) && input[i+count] == input[i] && count < 255 {
                count++
            }
            buf.WriteByte(byte(count))
            buf.WriteByte(input[i])
            i += count
        }

        return buf.Bytes()
    }
    ```

