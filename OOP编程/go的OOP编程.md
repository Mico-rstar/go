### OOP四大原则
- 封装
  - 将数据（属性）和操作数据的方法绑定在一起。
  - 控制对内部状态的访问（如私有/公有成员）
- 继承
  - 代码复用机制
- 多态
  - 编译时多态（方法重载）
  - 运行时多态（方法重写 + 动态绑定）同一接口可以有多种实现方式。
  - 同一接口可以有多种实现方式。
- 抽象
  - 隐藏复杂实现细节，暴露简洁接口。

### GO如何实现封装？
#### 封装属性和方法
go语言使用 **方法** ，通过为函数指定默认接收器来为类型设计专属函数，从而将属性和操作绑定在一起
```go
type Animal type {
    Name    string
}
func (a Animal) Speak() {...}
```
#### 控制可见性
Go语言只有一种控制可见性的规则：大写首字母的标识符会从定义它们的包中导出，小写字母的则不会
```go
// n在包之外不可见
type Counter struct { n int }
// 以下方法以大写字母开头，对包外可见
func (c *Counter) N() int     { return c.n }
func (c *Counter) Increment() { c.n++ }
func (c *Counter) Reset()     { c.n = 0 }

```

### GO如何实现类似继承的机制？
Go中没有继承机制，但继承本质是为了实现代码复用，而go的哲学认为 **组合优于继承**
Go中通过嵌入结构体来实现组合
#### 通过命名字段间接引入
此时复合类型只能间接调用内嵌类型的方法
```go
// 基础组件
type Camera struct {
    Megapixels int
    Brand      string
}

func (c Camera) TakePhoto() string {
    return fmt.Sprintf("拍摄了%d万像素的照片", c.Megapixels)
}

type Battery struct {
    Capacity int    // mAh
    Level    int    // 百分比
}

func (b *Battery) Charge() {
    b.Level = 100
    fmt.Println("电池充满电！")
}

func (b *Battery) Use(amount int) {
    b.Level -= amount
    if b.Level < 0 {
        b.Level = 0
    }
}

type Screen struct {
    Size   float64 // 英寸
    Type   string  // LCD, OLED
}

func (s Screen) Display(content string) {
    fmt.Printf("在%.1f英寸%s屏幕上显示：%s\n", s.Size, s.Type, content)
}

// 通过组合创建智能手机
type Smartphone struct {
    Camera  Camera  // 组合：手机有一个摄像头
    Battery Battery // 组合：手机有一个电池
    Screen  Screen  // 组合：手机有一个屏幕
    Brand   string
    Model   string
}
```
```go
// Smartphone的方法
func (p Smartphone) MakeCall(number string) {
    p.Battery.Use(1) // 使用电池
    fmt.Printf("正在拨打：%s\n", number)
}

func (p Smartphone) TakePicture() {
    photo := p.Camera.TakePhoto() // 委托给摄像头组件
    p.Battery.Use(2)
    fmt.Println("照片结果:", photo)
}

func (p Smartphone) ShowStatus() {
    fmt.Printf("手机：%s %s\n", p.Brand, p.Model)
    fmt.Printf("电池电量：%d%%\n", p.Battery.Level)
    p.Screen.Display("手机状态界面")
}
```

#### 通过匿名字段引入
通过匿名字段组合类型，复合类型将能够直接调用嵌套类型的方法
```go
type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64{
	return math.Sqrt(math.Pow(p.X - q.X, 2) + math.Pow(p.Y - q.Y, 2))
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
    *Point  // 这里用Point同样可以使用ScaleBy方法
    Color color.RGBA
}

p := ColoredPoint{&Point{1, 1}, red}
q := ColoredPoint{&Point{5, 4}, blue}
fmt.Println(p.Distance(*q.Point)) // "5"
p.ScaleBy(2) // 2, 2

```


### Go如何实现多态
#### interface的底层机制
[[interface的本质]]

#### 使用interface实现动态分发的多态特性
```go
type Speaker interface {
    Speak() string
}

type Cat struct{}
func (c Cat) Speak() string { return "Miao~" }

type Dog struct{ }
func (d Dog) Speak() string { return "Woof!" }

func MakeSound(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    var s Speaker = Dog{} // Dog 实现了 Speaker
    MakeSound(s)          // 输出: Woof!
    s = Cat{}
    MakeSound(s)          // 输出: Miao~
}
```

#### 重写方法特性
go中并不会出现真正的方法重写
```go
type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64{
	fmt.Println("Point.Distance")
	return math.Sqrt(math.Pow(p.X - q.X, 2) + math.Pow(p.Y - q.Y, 2))
}
type ColoredPoint struct {
    Point
}
// 尝试重写Distance方法
func (p ColoredPoint) Distance(q ColoredPoint) float64 {
	fmt.Println("ColoredPoint.Distance")
	return math.Sqrt(math.Pow(p.X - q.X, 2) + math.Pow(p.Y - q.Y, 2))
}
```
尝试调用Distance
```go
cp := ColoredPoint{Point{1.0, 2.0}}
qp := ColoredPoint{Point{1.3, 2.2}}
cp.Distance(qp)     // 调用的是 ColoredPoint.Distance
cp.Point.Distance(qp.Point)     // 通过显式调用依然可以调用Point.Distance, 说明没有发生真正的重写
```




### Go如何实现抽象
1. 使用接口隐藏具体实现
2. 使用大小写可见性控制来进程封装
3. 没有类，但使用结构体进行抽象，更接近本质