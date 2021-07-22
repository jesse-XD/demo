package closure

import "fmt"

// 闭包练习
// 闭包的原理
// -- 1.函数可以作为返回值
// -- 2.函数内部查找变量的顺序,先从内部查找，找不到再从外部查找

// 练习1
// 已知存在 f1、f2 函数，需要实现f1调用f2 f1(f2)
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 利用闭包的机制实现
// 定义一个闭包函数
func f3(x, y int) func() {
	return func() {
		f2(x, y)
	}
}

func F4() {
	ret := f3(1, 2)
	f1(ret)
}

// 闭包实现 2.0
func f3_2(f func(int, int), x, y int) func() {
	return func() {
		f(x, y)
	}
}

func F4_2() {
	ret := f3_2(f2, 10, 20)
	f1(ret)
}
