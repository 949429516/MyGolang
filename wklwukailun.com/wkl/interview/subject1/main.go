package main

/*
在A和B出填入代码，使输出为foo

type S struct {
	m string
}

func f() *S {
	return _____A
}

func main() {
	p := _____B
	fmt.Println(p.m)
}
*/
import "fmt"

type S struct {
	m string
}

func f() *S {
	return &S{m: "foo"}
}

func main() {
	p := f()
	fmt.Println(p.m)
}
