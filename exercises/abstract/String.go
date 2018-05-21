package main

import "fmt"

func main() {
	//建立一个go程序  实现累加和功能
	str := "A"
	for i := 0; i < 100; i++ {
		fmt.Printf("%s\n", str)
		str = str + "A"
	}
}
