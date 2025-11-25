package main

import (
	"fmt"
	"notes/src/method/counter"
	"notes/src/method/point"
	"image/color"
)

type Animal struct {
	name string
	sex  bool
	life int
}

type Food struct {
	name   string
	kaluli int
}

func (a *Animal) eat(food Food) {
	fmt.Println(a.name, " eat ", food.name)
}

type Color int 
const (
	White Color = iota
	Yellow
	Black
)


// 组合而非继承
type Rabbit struct {
	Animal
	color Color
}

func (r Rabbit) jump() {
	fmt.Printf("A rabbit named %s jumped, its color is %v\n", r.name, r.color)
}

func (r *Rabbit) sleep() {
	fmt.Println("Rabbit ", r.name, " is sleeping")
}

type A struct {
	x	int
}

type B struct {
	x	int
}

type C struct {
	a 	A
	B
}


func main() {
	animal := Animal{
		name: "chicken",
		sex: true,
		life: 5,
	}
	rice := Food {
		name: "rice",
		kaluli: 100,
	}
	animal.eat(rice)

	(&Animal{"rabbit", false, 6}).eat(rice)

	rabbit := Rabbit{
		Animal: Animal{
			name: "xiangsi",
			sex: false,
			life: 3,
		},
		color: White,
	}

	rabbit.eat(rice)
	rabbit.jump()

	c := C{
		a: A{1},
		B: B{2},
	}
	fmt.Println("c = ", c.a.x)
	rabbit.sleep()
	
	// jump := (*Rabbit).sleep

	NilTest()

	cnt := counter.New(1)
	fmt.Println(cnt.N())

	cp := point.ColoredPoint{
		Point: point.Point{1, 2},
		Color: color.RGBA{},
	}

	cp.ScaleBy(2)
	fmt.Println(cp)

	qp := point.ColoredPoint{
		Point: point.Point{3, 4},
		Color: color.RGBA{},
	}

	// cp.Distance(qp)
	// 显式调用子类型方法
	// cp.Point.Distance(qp.Point)
	cp.Distance(qp)
}