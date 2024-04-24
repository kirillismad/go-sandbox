package examples

import (
	"fmt"
	"math"
)

func DemoInterface() {
	s := []Geometry{
		newRectangle(4, 3),
		newRectangle(2, 3),
		newCircle(4),
		newCircle(5),
	}

	fmt.Println("Result area:", sumArea(s...))

	resize(1.2, s...)
	for _, g := range s {
		fmt.Println(g)
	}

	demoTypeAssertion()
}

func sumArea(gs ...Geometry) float64 {
	sum := 0.0
	for _, g := range gs {
		fmt.Println(g, "area", g.area())
		sum = sum + g.area()
	}
	return sum
}

func resize(k float64, gs ...Geometry) {
	for _, g := range gs {
		g.resize(k)
	}
}

type Geometry interface {
	perim() float64
	area() float64

	resize(k float64)
}

type Rectangle struct {
	height float64
	width  float64
}

func newRectangle(height, width float64) *Rectangle {
	r := new(Rectangle)
	r.height = height
	r.width = width
	return r
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(height: %v, width: %v)", r.height, r.width)
}

func (r Rectangle) perim() float64 {
	return r.height*2 + r.width*2
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (r *Rectangle) resize(k float64) {
	r.height = r.height * k
	r.width = r.width * k
}

type Circle struct {
	radius float64
}

func newCircle(radius float64) *Circle {
	c := new(Circle)
	c.radius = radius
	return c
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(radius: %v)", c.radius)
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r *Circle) resize(k float64) {
	r.radius = r.radius * k
}

type someInterface interface {
	method1() int
}

type someInterfaceImpl struct {
	field int
}

func (o someInterfaceImpl) method1() int {
	return o.field
}

type nestedInterface struct {
	someInterface
	name string
}

type IHandler interface {
	Handle(a, b int) int
}

type BaseHandler struct{}

func (h *BaseHandler) Handle(a, b int) int {
	fmt.Println(a, b)
	return 0
}

type SumHandler struct {
	base BaseHandler
}

func (h *SumHandler) Handle(a, b int) int {
	h.base.Handle(a, b)
	return a + b
}

type SubHandler struct{}

func (h *SubHandler) Handle(a, b int) int {
	return a - b
}

func TestHandlers() {
	sub := new(SubHandler)
	sum := &SumHandler{base: BaseHandler{}}

	handlers := make([]IHandler, 2)

	handlers[0] = sub
	handlers[1] = sum

	for i, handler := range handlers {
		fmt.Println(i, ")", handler.Handle(10, 2))
	}
}

func demoTypeAssertion() {
	var x interface{} = int(42)
	x1, ok := x.(int8)
	fmt.Println(x1, ok)
}
