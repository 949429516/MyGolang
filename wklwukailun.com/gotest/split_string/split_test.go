package splitstring

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := SplitString("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

/*
go test
PASS
ok      gotest/split_string     0.001s
*/
func TestSplitWithComplexSep(t *testing.T) {
	got := SplitString("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

/*
go test -v
=== RUN   TestSplit
--- PASS: TestSplit (0.00s)
=== RUN   TestSplitWithComplexSep
--- PASS: TestSplitWithComplexSep (0.00s)
PASS
ok      gotest/split_string     0.001s
*/

/*
只运行一个
go test -run=TestSplitWithComplexSep -v
*/

//子测试，多个用例并行
func TestSplitAll(t *testing.T) {
	// 定义测试表格
	// 这里使用匿名结构体定义了若干个测试用例
	// 并且为每个测试用例设置了一个名称
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"case1", "a:b:c", ":", []string{"a", "b", "c"}},
		{"case2", "a:b:c", ",", []string{"a:b:c"}},
		{"case3", "abcd", "bc", []string{"a", "d"}},
	}
	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
			t.Parallel() // 将每个测试用例标记为能够彼此并行运行
			got := SplitString(tt.input, tt.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected:%#v, got:%#v", tt.want, got)
			}
		})
	}
}

/*
go test -bench=Split  基准性能
go test -bench=Split -benchmem  内存
*/
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitString1("a:b:c:d:e", ":")
	}
}

/*
go test -bench=Fib1
可以指定cpu数量-cpu=1
goos: linux
goarch: amd64
pkg: gotest/split_string
cpu: AMD Ryzen 5 3500X 6-Core Processor
BenchmarkFib1-4         562376505                2.111 ns/op
BenchmarkFib10-4         4157536               284.9 ns/op
PASS
ok      gotest/split_string     2.888s
*/
/*
性能比较函数
*/
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}
func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
