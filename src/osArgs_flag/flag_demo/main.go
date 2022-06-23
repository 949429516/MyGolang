package main

import (
	"flag"
	"fmt"
)

func main() {
	//创建标志位参数
	name := flag.String("name", "王森堡", "请输入姓名")
	//age := flag.Int("age", 9000, "请输入年龄")
	//方法二
	var age int
	flag.IntVar(&age, "age", 9000, "请输入年龄")

	married := flag.Bool("married", false, "婚否")
	//使用flag
	flag.Parse()
	fmt.Println(*name, age, *married)

	fmt.Println(flag.Args())  //返回命令行后的其他参数[]string返回
	fmt.Println(flag.NArg())  //返回个数
	fmt.Println(flag.NFlag()) //返回命令行参数
}
