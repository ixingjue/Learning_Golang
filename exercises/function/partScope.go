package main

var a = 6
//在函数执行的时候 局部变量作用域  局部变量覆盖全局作用域
func main() {
	p()
	q()
	p()
}
func p() {
	println(a)
}
func q() {
	a := 5 //定义
	println(a)
}
