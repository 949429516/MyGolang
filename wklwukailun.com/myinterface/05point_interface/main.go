package main

import "fmt"

type cat struct {
	name string
	feet int
}
type animal interface {
	move()
	eat(string)
}

//指针接收者实接口现接口只能接收指针类型
// func (c *cat) move() {
// 	fmt.Println("走猫步......")
// }
// func (c *cat) eat(food string) {
// 	fmt.Printf("猫吃%s\n", food)
// }

//值接受者实现接口，即可以接收值类型也可以接收指针类型
func (c cat) move() {
	fmt.Println("走猫步......")
}
func (c cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}
func main() {
	var a1 animal
	c1 := cat{"tom", 4} //cat
	c2 := &cat{"二蛋", 4} //*cat
	a1 = c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
