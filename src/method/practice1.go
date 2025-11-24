package main

// import "fmt"

type Calculator struct {
    value float64
}

// TODO: 为Calculator添加Add方法，接收一个float64参数
func (c *Calculator) Add(num float64) {
    c.value += num
}

// TODO: 为Calculator添加Multiply方法，接收一个float64参数
func (c *Calculator) Multiply(factor float64) {
    c.value *= factor
}

// func main() {
//     calc := &Calculator{value: 10}

//     // TODO: 将Add方法创建为方法值
//     addOp := calc.Add  // 方法值：已经绑定了calc对象

//     // TODO: 将Multiply方法创建为方法值
//     multiplyOp := calc.Multiply  // 方法值：已经绑定了calc对象

//     // 使用方法值进行计算
//     // 要求：addOp(calc, 5) 应该得到 15
//     // 要求：multiplyOp(calc, 2) 应该得到 30

//     fmt.Printf("初始值: %.2f\n", calc.value) // 10.00

//     // 使用方法值 - 注意方法值已经绑定了calc，只需要传参数
//     addOp(5)          // calc.value = 10 + 5 = 15
//     fmt.Printf("加法后: %.2f\n", calc.value) // 15.00

//     multiplyOp(2)     // calc.value = 15 * 2 = 30
//     fmt.Printf("乘法后: %.2f\n", calc.value) // 30.00

//     fmt.Printf("最终结果: %.2f\n", calc.value)
// }