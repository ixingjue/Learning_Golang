package main

import "fmt"

func main() {
	var p *int
	//打印空指针
	fmt.Printf("%v\n", p)
	var i int
	p = &i //获取i的地址
	//从指针获取值是通过在指针变量前置’*’ 实现的：
	*p = 8 //修改i的值
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", *p)
	fmt.Printf("%v", i)
}
