package main

import "fmt"

/*GO语言中函数的return不是原子操作，在底层是分为两步来执行
1.返回值赋值
defer
2.真正的RET返回
函数中如果存在defer,那么defer执行的时机是在1与2之间
*/

func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x不是返回值
	}()
	return x
}

//返回值在定义的时候已经是x
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值=x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ //defer修改的x
	}()
	return x //返回值=y=x=5
}

//return:
//第一步: 给返回变量x赋值  x = 5
//defer: 将x传入defer  是值传递  修改copy后的x的值 x` = 6  x = 5
//第二步: 返回变量x  x = 5
//输出: 5
func f4() (x int) {
	defer func(x int) {
		x++
	}(x) //这里相当于复制了一个,在不同作用域
	return 5
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
