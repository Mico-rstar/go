package point

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64{
	fmt.Println("Point.Distance")
	return math.Sqrt(math.Pow(p.X - q.X, 2) + math.Pow(p.Y - q.Y, 2))
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
    Point
    Color color.RGBA
}

// 尝试重写Distance方法
func (p ColoredPoint) Distance(q ColoredPoint) float64 {
	fmt.Println("ColoredPoint.Distance")
	return math.Sqrt(math.Pow(p.X - q.X, 2) + math.Pow(p.Y - q.Y, 2))
}

