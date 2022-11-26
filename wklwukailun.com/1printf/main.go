package main

import "fmt"

type simple struct {
	value int
}

func main() {
	a := simple{
		value: 10,
	}
	// 通用占为符
	fmt.Println("===================通用占为符===================")
	fmt.Printf("默认格式的值,%v\n", a)
	fmt.Printf("版汉字段名的值,%+v\n", a)
	fmt.Printf("go语法表示的值,%#v\n", a)
	fmt.Printf("go语法表示的类型,%T\n", a)
	fmt.Printf("输出字面的百分号,%%10\n")
	// 整数类型占位符
	fmt.Println("===================整数类型占位符===================")
	v1 := 10
	v2 := 20170
	fmt.Printf("二进制,%b \n", v1)
	fmt.Printf("Unicode码点转字符,%c \n", v2)
	fmt.Printf("十进制,%b \n", v1)
	fmt.Printf("八进制,%b \n", v1)
	fmt.Printf("Oo为前缀的八进制,%o \n", v1)
	fmt.Printf("单引号将字符值包起来,%q \n", v2)
	fmt.Printf("十六进制,%x \n", v1)
	fmt.Printf("大写十六进制,%X \n", v1)
	fmt.Printf("Unicode格式,%U \n", v2)
	//宽度设置
	fmt.Println("===================宽度设置===================")
	fmt.Printf("%v的二进制为:%b; go语法表示的二进制为:%#b; 指定二进制宽度为8位,不是8位的补0:%08b \n", v1, v1, v1, v1)
	fmt.Printf("%v的十六进制为:%x; 使用go语法表示,宽度为8位,不足8位的补0:%#08x\n", v1, v1, v1)
	fmt.Printf("%v的字符为:%c; 指定宽度为3位,不足5位补空格:%5c \n", v2, v2, v2)
	//浮点类型
	fmt.Println("===================浮点类型===================")
	var f1 float64 = 123.789
	var f2 float64 = 12345678910.78999
	fmt.Printf("指数为二的幂的无小数科学计数法:%b \n", f1)
	fmt.Printf("科学计数法:%e \n", f1)
	fmt.Printf("科学计数法大写:%E \n", f1)
	fmt.Printf("有小数点而无指数,即常规的浮点数格式默认宽度和精度:%f \n", f1)
	fmt.Printf("宽度为9精度默认:%9f \n", f1)
	fmt.Printf("默认宽度,精度保留2位小数:%.2f \n", f1)
	fmt.Printf("根据情况自动选择%%e 或 %%f 未输出,产生的更紧凑的输出(末尾无0):%g %g \n", f2, f2)
	fmt.Printf("根据情况自动选择%%E 或 %%f 未输出,产生的更紧凑的输出(末尾无0):%G %G \n", f2, f2)
	fmt.Printf("以十六进制方式表示: %x \n", f1)
	fmt.Printf("以十六进制方式表示,大写: %X \n", f1)
	//字符串
	fmt.Println("===================字符串===================")
	var str = "今天是个好日子"
	fmt.Printf("%s \n", str)
	fmt.Printf("%q \n", str)
	fmt.Printf("字符串16进制表示,每两个字母为一个byte%x \n", str)
	fmt.Printf("字符串16进制表示,以空格分割,每两个字母为一个byte% X \n", str)
}
