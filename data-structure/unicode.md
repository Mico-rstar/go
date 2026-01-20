# ch3/unicode

#### unicode
1. 字符属性判断
    ```go
    // 判断字符类型
    unicode.IsLetter('A')        // true (字母)
    unicode.IsDigit('5')         // true (数字)
    unicode.IsSpace(' ')         // true (空白字符)
    unicode.IsPunct('.')         // true (标点符号)
    unicode.IsSymbol('$')        // true (符号)
    unicode.IsMark('́')          // true (组合字符)

    // 判断大小写
    unicode.IsUpper('A')         // true (大写)
    unicode.IsLower('a')         // true (小写)
    unicode.IsTitle('ǅ')         // true (标题格式)
    ```

2. 字符转换
    ```go
    // 大小写转换
    unicode.ToLower('A')         // 'a'
    unicode.ToUpper('a')         // 'A'
    unicode.ToTitle('a')         // 'A' (标题格式)

    // 转换为特定格式
    unicode.To(unicode.UpperCase, 'a')   // 'A'
    unicode.To(unicode.LowerCase, 'A')   // 'a'
    unicode.To(unicode.TitleCase, 'a')   // 'A'
    ```

3. 脚本和范围判断
    ```go
    // 判断字符所属脚本
    unicode.Is(unicode.Han, '中')           // true (汉字)
    unicode.Is(unicode.Latin, 'A')          // true (拉丁字母)
    unicode.Is(unicode.Cyrillic, 'А')       // true (西里尔字母)
    unicode.Is(unicode.Arabic, 'ا')         // true (阿拉伯字母)

    // 判断字符范围
    unicode.In('中', unicode.Han)           // true
    unicode.In('A', unicode.Latin)          // true
    ```

4. 特殊字符类别
    ```go
    // 控制字符
    unicode.IsControl('\n')      // true
    unicode.IsControl('\t')      // true

    // 数字相关
    unicode.IsNumber('₃')        // true (数字)
    unicode.IsDigit('3')         // true (十进制数字)

    // 标点符号
    unicode.IsGraphic('!')       // true
    unicode.IsPrint('!')         // true
    unicode.IsGraphic('\n')      // false
    unicode.IsPrint('\n')        // false
    ```

5. 字符分类表
    ```go
    // 使用分类表
    var greek = unicode.RangeTable{
        R16: []unicode.Range16{{0x370, 0x3ff, 1}, {0x1f00, 0x1fff, 1}},
    }
    unicode.Is(&greek, 'α')      // true

    // 预定义分类
    unicode.Is(unicode.Greek, 'α')     // true
    unicode.Is(unicode.Hiragana, 'あ') // true
    unicode.Is(unicode.Katakana, 'カ') // true
    ```

6. 简单折叠
    ```go
    // 大小写折叠（不区分大小写的比较）
    unicode.SimpleFold('A')     // 'a'
    unicode.SimpleFold('a')     // 'A'
    unicode.SimpleFold('İ')     // 'i̇' (土耳其大写I)

    // 折叠迭代
    r := 'A'
    for i := 0; i < 3; i++ {
        r = unicode.SimpleFold(r)
    }
    // r会遍历该字符的所有大小写形式
    ```

7. UTF-8相关
    ```go
    // UTF-8编码长度
    utf8.RuneLen('A')           // 1
    utf8.RuneLen('中')          // 3
    utf8.RuneLen('€')           // 3

    // 解码UTF-8
    data := []byte("hello")
    r, size := utf8.DecodeRune(data)  // 'h', 1
    ```

8. 字符集和编码
    ```go
    // UTF-16代理对判断
    utf8.IsSurrogate(rune(0xd800))    // true (高代理)
    utf8.IsSurrogate(rune(0xdc00))    // true (低代理)

    // Rune count
    utf8.RuneCountInString("hello中文") // 7
    utf8.RuneCount([]byte("hello中文")) // 7
    ```

9. 特殊Unicode属性
    ```go
    // 组合字符
    unicode.IsMark('́')          // true (组合重音)
    unicode.IsMark('\u0301')     // true (组合重音)

    // 变体选择器
    unicode.IsVariantSelector(rune(0xfe00))  // true

    // 非字符
    unicode.IsNoncharacter(rune(0xfffe))    // true
    ```

10. 实际应用示例
    ```go
    // 判断是否为有效标识符字符
    func isValidIdentifier(r rune) bool {
        return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
    }

    // 文本统计分析
    func analyzeText(text string) {
        letterCount := 0
        digitCount := 0
        spaceCount := 0
        chineseCount := 0

        for _, r := range text {
            switch {
            case unicode.IsLetter(r):
                letterCount++
                if unicode.Is(unicode.Han, r) {
                    chineseCount++
                }
            case unicode.IsDigit(r):
                digitCount++
            case unicode.IsSpace(r):
                spaceCount++
            }
        }

        fmt.Printf("字母: %d (含中文: %d), 数字: %d, 空白: %d\n",
            letterCount, chineseCount, digitCount, spaceCount)
    }

    // 大小写不敏感搜索
    func caseInsensitiveSearch(text, pattern string) bool {
        textRunes := []rune(text)
        patternRunes := []rune(pattern)

        if len(textRunes) < len(patternRunes) {
            return false
        }

        for i := 0; i <= len(textRunes)-len(patternRunes); i++ {
            match := true
            for j := 0; j < len(patternRunes); j++ {
                if !caseEqual(textRunes[i+j], patternRunes[j]) {
                    match = false
                    break
                }
            }
            if match {
                return true
            }
        }
        return false
    }

    func caseEqual(a, b rune) bool {
        // 考虑所有大小写形式
        for ar := a; ar <= b; ar = unicode.SimpleFold(ar) {
            for br := b; br <= a; br = unicode.SimpleFold(br) {
                if ar == br {
                    return true
                }
            }
        }
        return false
    }
    ```

