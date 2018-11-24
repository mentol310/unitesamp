package main

import (
	"reflect"
	"testing"
)

func TestPointer(t *testing.T) {
	expect := 6

	var i int
	p := &i
	i = 5
	*p++
	actual := i

	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}
}

func TestPointerPow(t *testing.T) {
	pow := func(p *[3]int) {
		i := 0
		for i < 3 {
			p[i] = p[i] * p[i]
			i++
		}
	}

	p := &[3]int{1, 2, 3}
	pow(p)

	expect := [3]int{1, 4, 9}
	actual := *p

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}
