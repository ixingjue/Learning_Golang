package main

import "fmt"
//扩展/修改上一个问题的程序，替换位置 4 开始的三个字符为 “abc”
func main() {
	s := "helloworld"
	r := []rune(s)
	copy(r[4:4+3], []rune("abc"))
	fmt.Printf("Before: %s\n", s)
	fmt.Printf("After : %s\n", string(r))
}
