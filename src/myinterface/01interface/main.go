package main

/*
接口是一种类型，一种抽象的类型
接口就是要实现方法的清单
*/
import "fmt"

type cat struct {
	name string
	feet int
}

type animal interface {
	move()
	eat(string)
}

func (c cat) move() {
	fmt.Println("走猫步")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func main() {
	var b animal

	a := cat{
		name: "波斯猫",
		feet: 4,
	}
	b = a
	b.move()
	b.eat("大便")
}
