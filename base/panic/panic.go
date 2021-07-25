package panic

import "fmt"

func FuncTest(){
	funcA()
	funcB()
	funcC()
}

func funcA(){
	fmt.Println("A")
}

func funcB(){
	defer func(){
		// 捕捉到程序崩溃的信息
		err := recover()
		fmt.Println("程序出错：", err)
	}()

	fmt.Println("B")
	panic("程序崩溃了。。。")
}
func funcC(){
	fmt.Println("C")
}