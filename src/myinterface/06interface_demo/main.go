package main

import "fmt"

type speaker interface {
	speak()
}
type cat struct{}
type dog struct{}
type person struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}
func (d dog) speak() {
	fmt.Println("汪汪王~")
}
func (d person) speak() {
	fmt.Println("啊啊啊~")
}

func da(x speaker) {
	x.speak()
}

func main() {
	c := cat{}
	d := dog{}
	p := person{}
	da(c)
	da(d)
	da(p)
}
