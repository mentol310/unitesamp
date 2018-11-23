package main

import (
	"reflect"
	"testing"
)

func TestMakeMap(t *testing.T) {
	expect := make(map[int]string)
	expect[1] = "itti"
	expect[3] = "san"
	expect[5] = "go"

	actual := map[int]string{
		1: "itti",
		3: "san",
		5: "go",
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestMakeSliceMap(t *testing.T) {
	expect := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: []int{1, 2, 3},
	}
	actual := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestMapGetter(t *testing.T) {
	base := map[int]string{1: "A", 2: "B", 3: "C"}
	expect := "B"
	actual := base[2]
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestMapFor(t *testing.T) {
	base := map[int]string{
		1: "Apple",
	}
	expectIndex := 1
	expectValue := "Apple"

	for k, v := range base {
		if expectIndex != k {
			t.Errorf("%v != %v", expectIndex, k)
		}
		if expectValue != v {
			t.Errorf("%s != %s", expectValue, v)
		}
	}
}

func TestMapLen(t *testing.T) {
	base := map[int]string{1: "A", 2: "B", 3: "c"}
	expect := 3
	actual := len(base)
	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestMapDelete(t *testing.T) {
	expect := map[int]string{1: "a", 3: "c"}
	actual := map[int]string{1: "a", 2: "b", 3: "c"}
	delete(actual, 2)
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}
