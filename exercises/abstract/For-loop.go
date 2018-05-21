package main

import "fmt"

func main() {
	//方法1  创建一个基于 for 的简单的循环。使其循环 10 次，并且使用 fmt 包打印出计数器的值。
	/*for i := 0;i < 10; i++{
	fmt.Printf("%d\n",i)
	}*/
	//方法2  用 goto 改写 1 的循环。关键字 for 不可使用。
	/*i := 0
Loop:
	fmt.Printf("%d\n", i)
	if i < 10 {
		i++
		goto Loop
	}*/
	//方法3  再次改写这个循环，使其遍历一个 array，并将这个 array 打印到屏幕上。
	/*var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Printf("%v", arr)*/
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%v\n", a)

}
