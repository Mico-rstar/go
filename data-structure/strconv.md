# ch3/strconv

#### strconv
1. 字符串转数值类型
    ```go
    // 字符串转整数
    strconv.Atoi("123")              // 123, nil
    strconv.ParseInt("123", 10, 64)  // 123, nil
    strconv.ParseUint("123", 10, 64) // 123, nil

    // 字符串转浮点数
    strconv.ParseFloat("123.45", 64) // 123.45, nil

    // 带错误的转换
    _, err := strconv.Atoi("abc")    // 0, "strconv.Atoi: parsing \"abc\": invalid syntax"
    ```

2. 数值类型转字符串
    ```go
    // 整数转字符串
    strconv.Itoa(123)           // "123"
    strconv.FormatInt(123, 10)  // "123"
    strconv.FormatUint(123, 10) // "123"

    // 浮点数转字符串
    strconv.FormatFloat(123.456, 'f', 2, 64) // "123.46"
    strconv.FormatFloat(123.456, 'e', -1, 64) // "1.23456e+02"

    // 通用格式化
    strconv.FormatBool(true)    // "true"
    ```

3. 布尔值转换
    ```go
    // 字符串转布尔值
    strconv.ParseBool("true")   // true, nil
    strconv.ParseBool("1")      // true, nil
    strconv.ParseBool("0")      // false, nil
    strconv.ParseBool("false")  // false, nil

    // 布尔值转字符串
    strconv.FormatBool(true)    // "true"
    strconv.FormatBool(false)   // "false"
    ```

4. 不同进制转换
    ```go
    // 十进制转其他进制字符串
    strconv.FormatInt(255, 2)   // "11111111" (二进制)
    strconv.FormatInt(255, 8)   // "377" (八进制)
    strconv.FormatInt(255, 16)  // "ff" (十六进制)
    strconv.FormatInt(255, 32)  // "7v" (三十二进制)

    // 其他进制转十进制
    strconv.ParseInt("ff", 16, 64)  // 255, nil
    strconv.ParseInt("11111111", 2, 64)  // 255, nil
    ```

5. 引号和转义
    ```go
    // 字符串转Go语法格式
    strconv.Quote("hello\nworld")      // "\"hello\\nworld\""
    strconv.QuoteRune('a')             // "'a'"
    strconv.QuoteToASCII("中文")        // "\"\\u4e2d\\u6587\""

    // Go语法格式转原始字符串
    strconv.Unquote("\"hello\\nworld\"") // "hello\nworld", nil
    ```

6. 追加到字节切片
    ```go
    var buf []byte
    buf = strconv.AppendInt(buf, 123, 10)     // []byte("123")
    buf = strconv.AppendFloat(buf, 123.456, 'f', 2, 64) // []byte("123123.46")
    buf = strconv.AppendBool(buf, true)        // []byte("123123.46true")
    buf = strconv.AppendQuote(buf, "hello")    // []byte("123123.46true\"hello\"")
    ```

7. 错误处理
    ```go
    // 检查数值范围
    _, err := strconv.ParseInt("999999999999999999999", 10, 64)
    // err: "strconv.ParseInt: parsing \"999999999999999999999\": value out of range"

    // 自定义错误处理
    if num, err := strconv.Atoi("123"); err != nil {
        log.Printf("转换失败: %v", err)
    } else {
        fmt.Printf("转换成功: %d", num)
    }
    ```

8. 特殊浮点值
    ```go
    // 解析特殊浮点值
    strconv.ParseFloat("NaN", 64)    // NaN, nil
    strconv.ParseFloat("+Inf", 64)   // +Inf, nil
    strconv.ParseFloat("-Inf", 64)   // -Inf, nil

    // 格式化特殊值
    strconv.FormatFloat(math.NaN(), 'g', -1, 64)  // "NaN"
    strconv.FormatFloat(math.Inf(1), 'g', -1, 64) // "+Inf"
    ```

9. 位大小控制
    ```go
    // 解析指定位大小的整数
    strconv.ParseInt("127", 10, 8)   // 127, nil (int8)
    strconv.ParseInt("128", 10, 8)   // 0, error (超出int8范围)
    strconv.ParseInt("-32768", 10, 16) // -32768, nil (int16)
    ```

10. 实用转换函数
    ```go
    // 字符串转字符
    rune, _, err := strconv.DecodeRuneInString("hello")  // 'h', 1, nil

    // 字符转字符串
    strconv.QuoteRune('中')  // "'中'"

    // 判断是否为可打印字符
    strconv.IsPrint('a')     // true
    strconv.IsPrint('\n')    // false
    ```

