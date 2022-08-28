package main

import "fmt"

type Options struct {
	opt1 string
	opt2 string
	opt3 string
	opt4 int
	opt5 int
	opt6 int
}

// 声明一个函数类型的变量，用于传参数
type Option func(opts *Options)

// 初始化结构体
func InitOptions(opts ...Option) (ret *Options) {
	ret = &Options{}
	for _, opt := range opts {
		// 调用函数传参
		opt(ret)
	}
	return
}
func main() {
	A := InitOptions(WithStrOption1("你"), WithStrOption2("好"), WithStrOption3("wsb"), WithStrOption4(1024))
	fmt.Println(A)
}

// 定义具体赋值的方法
func WithStrOption1(str string) Option {
	return func(opts *Options) {
		opts.opt1 = str
	}
}
func WithStrOption2(str string) Option {
	return func(opts *Options) {
		opts.opt2 = str
	}
}
func WithStrOption3(str string) Option {
	return func(opts *Options) {
		opts.opt3 = str
	}
}
func WithStrOption4(i int) Option {
	return func(opts *Options) {
		opts.opt4 = i
	}
}
func WithStrOption5(i int) Option {
	return func(opts *Options) {
		opts.opt5 = i
	}
}
func WithStrOption6(i int) Option {
	return func(opts *Options) {
		opts.opt6 = i
	}
}
