package shapes

import (
	"fmt"
	"math"
)

// Интерфейс Shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Прямоугольник
type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Круг
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Квадрат
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

// Треугольник
type Triangle struct {
	SideA, SideB, SideC float64
}

func (t Triangle) Area() float64 {
	// Формула Герона для площади треугольника
	s := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

// Вывод информации о фигуре
func PrintShapeDetails(s Shape) {
	fmt.Printf("Shape Details - Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}
