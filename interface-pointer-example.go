package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

type TP struct {
	S string
}

func (t T) M() {
	t.S = t.S + " added value"
	fmt.Println(t.S)
}

func (t *TP) M() {
	t.S = t.S + " added value"
	fmt.Println(t.S)
}

func main() {
	var i I

	i = T{"Hello"}
	i.M()
	describe(i)

	i = &TP{"Hello"}
	i.M()
	describe(i)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
