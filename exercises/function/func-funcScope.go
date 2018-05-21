package main
//局部变量仅仅在执行定义它的函数时有效。
var a int
func main() {
	a=5
	println(a)
	f()
}
func f()  {
	a:=6
	println(a)
	g()
}
func g()  {
	println(a)
}

