package main

import (
	"testing"
)

func TestChanel(t *testing.T) {
	ch := make(chan int, 10)
	ch <- 5

	expect := 5
	actual := <-ch

	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}
}
