### Tag
给结构体字段添加的元数据，可以通过反射机制来获取

```Go
type Action struct {
	Run		bool	`json:"run"`
	Swim	bool	`json:"swim"`
	Fly		bool 	`json:"fly"`
}

type Animal struct {
	Sex		uint8	`json:"sex"`
	Name	string	`json:"name"`
	Age 	int		`json:"age"`
	Action	Action	`jsonP:"action"`
}
```

```Go
type SyntaxExample struct {
    // 基本语法: `key:"value"`
    Field1 string `json:"name"`
    
    // 多个键值对: `key1:"value1" key2:"value2"`
    Field2 string `json:"age" xml:"user_age"`
    
    // 带选项的值: `key:"value,option1,option2"`
    Field3 int `json:"score,omitempty,string"`  // omitempty表示字段允许留空
    
    // 空标签: `key:""` 或 `key:","`
    Field4 bool `json:","`
    
    // 原始字符串: 标签内容按字面意义解析
    Field5 string `特殊 字符:"值 包含 空格"`
}
```


### 序列化
```Go
// 序列化
bird := Animal{
    Sex: 0,
    Name: "bird",
    Age: 2,
    Action: Action{
        Run: false,
        Swim: false,
        Fly: true,
    },
}
data, err = json.Marshal(bird)
if err != nil {
    log.Fatal("序列化错误：", err)
}
fmt.Println(string(data))
```
```Go
// 带缩进的序列化
data, err = json.MarshalIndent(bird, "", "  ")
```

### 反序列化
```Go
var decodeAnimal Animal 
err = json.Unmarshal(data, &decodeAnimal)
if err != nil {
    log.Fatal(err)
}
```



