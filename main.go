package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World !!")

	p := &[3]int{1, 2, 3}
	for k, v := range p {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
}
