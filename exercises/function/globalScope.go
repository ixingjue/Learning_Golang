package main

//全局变量作用域
var a = 6

func main() {
	p()
	q()
	p()
}
func p() {
	println(a)
}
func q() {
	a = 5 //赋值
	println(a)
}
