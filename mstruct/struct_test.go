package mstruct

import (
	"reflect"
	"testing"
)

func TestCreateStruct(t *testing.T) {
	pair1 := [2]int{1, 2}
	pair2 := IntPair{1, 2}

	expectLen := len(pair1)
	actualLen := len(pair2)

	if expectLen != actualLen {
		t.Errorf("%v != %v", expectLen, actualLen)
	}

	expectCap := cap(pair1)
	actualCap := cap(pair2)

	if expectCap != actualCap {
		t.Errorf("%v != %v", expectCap, actualCap)
	}

	expectValue := pair1[1]
	actualValue := pair2[1]

	if expectValue != actualValue {
		t.Errorf("%v != %v", expectValue, actualValue)
	}

	expect := pair1
	actual := pair2

	if reflect.DeepEqual(expect, actual) {
		t.Errorf("%v == %v", expect, actual)
	}
}

func TestCallbackType(t *testing.T) {
	Sum := func(ints []int, callback Callback) int {
		var sum int
		for _, v := range ints {
			sum += v
		}
		return callback(sum)
	}

	expect := 30
	actual := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)

	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}
}

func TestPointStruct(t *testing.T) {
	expect := Point{X: 10, Y: 5}

	var actual Point
	actual.X = 10
	actual.Y = 5

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestStructChild(t *testing.T) {
	bar := Bar{
		Hoge: Hoge{
			Foo: Foo{
				Name: "foofoo",
			},
		},
	}
	expect := "foofoo"
	actual := bar.Hoge.Foo.Name

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestNonameStruct(t *testing.T) {
	s := struct{ X int }{X: 11}

	expect := 11
	actual := s.X

	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}
}

func TestStructPointer(t *testing.T) {
	swap := func(p *Point) {
		x, y := p.Y, p.X
		p.X = x
		p.Y = y
	}

	expect := Point{X: 2, Y: 1}
	actual := Point{X: 1, Y: 2}
	swap(&actual)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestStructMethod(t *testing.T) {
	p := Point{X: 1, Y: 2}

	expect := "X: 1, Y: 2"
	actual := p.ToString()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestStructConstruct(t *testing.T) {
	expect := Point{1, 2}
	actual := NewPoint(1, 2)

	if !reflect.DeepEqual(expect, *actual) {
		t.Errorf("%v != %v", expect, *actual)
	}
}
