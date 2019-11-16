package calc_aera

import (
	"fmt"
	"math"
)

// 目标，实现五种图形的定义，包括：三角形，矩形，菱形，梯形，圆形
// 基于五种类型，需要求出分别的周长和面积
// 并且定义实现了 周长和面积的 即为实现了 AeraInterface

type Triangle struct {
	// 三角形
	a float64
	b float64
	c float64
}

type Rectangle struct {
	// 矩形
	width float64
	length float64
}

type Diamond struct {
	// 棱形
	length float64
	angle int
}

type Trapezoid struct {
	// 梯形
	top float64
	bottom float64
	height float64
}

type Circle struct {
	// 圆形
	radius float64
}

type AeraInterface interface {
	get_perimeter() float64
	get_aera() float64
	String() string
}

func (t Triangle) String() string {
	s := fmt.Sprintf("Triangle(a: %.2f, b: %.2f, c: %.2f)", t.a, t.b, t.c)
	return s
}

func (r Rectangle) String() string {
	s := fmt.Sprintf("Rectangle(length: %.2f, width: %.2f)", r.length, r.width)
	return s
}

func (d Diamond) String() string {
	s := fmt.Sprintf("Diamond(length: %.2f, angle: %d)", d.length, d.angle)
	return s
}

func (t Trapezoid) String() string {
	s := fmt.Sprintf("Trapezoid(top: %.2f, bottom: %.2f, height: %.2f)", t.top, t.bottom, t.height)
	return s
}

func (c Circle) String() string {
	s := fmt.Sprintf("Circle(radius: %.2f)", c.radius)
	return s
}


func (t Triangle) check_legal_triangle() bool {
	return t.a + t.b > t.c && t.b + t.c > t.a && t.a + t.c > t.b
}

func (t Triangle) get_perimeter() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) get_aera() float64 {
	perimeter := t.get_perimeter()
	half_perimeter := perimeter / 2
	aero := math.Sqrt(half_perimeter * (half_perimeter-t.a) * (half_perimeter - t.b) * (half_perimeter - t.c))
	return aero
}

func (r Rectangle) get_perimeter() float64 {
	return (r.length + r.width) * 2
}

func (r Rectangle) get_aera() float64 {
	return r.length * r.width
}

func (d Diamond) get_perimeter() float64 {
	return d.length * 4
}

func (d Diamond) get_aera() float64 {
	return math.Sin(float64(d.angle)) * d.length * d.length
}

func (t Trapezoid) get_perimeter() float64 {
	return 0
}

func (t Trapezoid) get_aera() float64 {
	return (t.top + t.bottom) * t.height / 2
}

func (c Circle) get_perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) get_aera() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func Run_calc_aera() {
	shape_slice := make([]AeraInterface, 5)
	triangle := Triangle{3,4,5}
	rectangle := Rectangle{3,4}
	diamond := Diamond{2,33}
	trapezoid := Trapezoid{2,3,5}
	circle := Circle{4}
	shape_slice[0], shape_slice[1], shape_slice[2], shape_slice[3], shape_slice[4] = triangle, rectangle, diamond, trapezoid, circle
	for _, value := range shape_slice {
		per := value.get_perimeter()
		aera := value.get_aera()
		fmt.Printf("%v 的 周长是：%f，面积是：%f\n", value, per, aera)
	}
}
