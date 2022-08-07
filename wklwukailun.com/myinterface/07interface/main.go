package main

import "fmt"

type anaimal interface {
	move()
	speak()
}

func task() anaimal {
	return newcat()
}

type cat struct {
	feet int
	name string
}

func newcat() cat {
	return cat{
		feet: 4,
		name: "李沁演",
	}
}
func (c cat) move() {
	fmt.Println("猫用", c.feet, "只脚跑路")
}
func (c cat) speak() {
	fmt.Println(c.name, "说日尼玛王森堡")
}

func main() {
	an := task()
	an.move()
	an.speak()
}
