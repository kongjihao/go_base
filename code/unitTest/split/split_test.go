package split

// split/split_test.go

import (
	"fmt"
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
// 	got := Split("a:b:c", ":")         // 程序输出的结果
// 	want := []string{"a", "b", "c"}    // 期望的结果
// 	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
// 		t.Errorf("expected:%v, got:%v", want, got) // 测试失败输出错误提示
// 	}
// }

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// 使用测试组的方式测试
// go test -v
// go test -v -run="More"
// go test -cover
func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

// 基准测试代码
// go test -bench=Split -benchmem
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

// fib_test.go
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

// 性能比较函数
//
//	go test -bench=.
func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }

// 实例函数的测试
// go test -run Example
func ExampleSplit() {
	fmt.Println(Split("a:b:c", ":"))
	fmt.Println(Split("沙河有沙又有河", "沙"))
	// Output:
	// [a b c]
	// [ 河有 又有河]
}
