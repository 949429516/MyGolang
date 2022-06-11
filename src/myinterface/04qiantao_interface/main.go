package main

import "fmt"

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int
}

func (c cat) move() {
	fmt.Println("走猫步......")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func main() {
	var a animal
	c := cat{"波斯猫", 4}
	a = c
	a.move()
	a.eat("大便")
}
