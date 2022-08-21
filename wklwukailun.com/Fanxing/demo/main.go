package main

import (
	"fmt"
)

// 泛型函数
func sum[T int | float64 | string](x, y T) bool {
	return x == y
}

// 泛型结构体
type Person[T any] struct {
	Name string
	Sex  T
}
type Personmany[T string | int | float64] struct {
	Name string
	Sex  T
}

// 泛型map
type TMap[k string | int, v any] map[k]v

// comparable可比较类型
type TMap2[k comparable, v any] map[k]v

// 泛型slice
type TSlice[s any] []s

// 泛型约束
type MyType interface{ string | int | float64 | bool }

func typeTest[T MyType](s T) {
	fmt.Println(s)
}

// ~用法
type myT interface {
	// 在int范围内的都可被约束
	~int | ~string
}
type Myint int

func test[T myT](t T) {
	fmt.Println(t)
}
func main() {
	fmt.Println(sum[int](1, 1))
	fmt.Println(sum[float64](1.1, 1.2))
	fmt.Println(sum[string]("1", "1"))
	fmt.Println(sum("1", "1"))
	// 泛型结构体
	p := Person[string]{Name: "王森堡", Sex: "男性"}
	p1 := Person[int]{Name: "王森堡", Sex: 0}
	p2 := Person[any]{Name: "王森堡", Sex: 0}
	p3 := Personmany[string]{Name: "王森堡", Sex: "男性"}
	fmt.Println(p, p1, p2, p3)
	// map类型
	m := make(TMap2[int, string], 0)
	m[1] = "一"
	fmt.Println(m)
	// 泛型slice
	s := TSlice[int]{}
	fmt.Printf("类型:%T\n", s)
	fmt.Println(s)
	// 泛型约束
	typeTest(1)

	test(123)
	test[Myint](456)
}
