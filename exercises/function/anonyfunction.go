package main

import "fmt"

func main() {
	abc:= func() {
		println("hello")
	}
	//调用匿名函数
	abc()
	fmt.Printf("%T\n",abc)
}
