package mstruct

import "fmt"

type IntPair [2]int

type Point struct {
	X, Y int
}

func NewPoint(x, y int) *Point {
	p := new(Point)
	p.X = x
	p.Y = y
	return p
}

func (p *Point) ToString() string {
	return fmt.Sprintf("X: %d, Y: %d", p.X, p.Y)
}

type Callback func(i int) int

type Foo struct {
	Name string
}
type Hoge struct {
	Foo
}
type Bar struct {
	Hoge
}
