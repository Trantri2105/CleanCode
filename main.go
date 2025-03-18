package main

import "fmt"

type A struct {
}

func (A) saySomething() {
	fmt.Println("hello A")
}

type B struct {
	A
}

func main() {
	b := B{}
	b.saySomething()
}
