package main

import "fmt"

func main() {
	
}
func ReadWrite() bool  {
	file.Open("file")
	defer file.Close()  //file.Close被添加到defer列表
	//做一些工作
	if failureX {
		return false //Close() 现在自动调用
	}
	if failureY {
		return false //这里也是
	}
	return true //And here

	for i := 0 ; i < 5 ; i++ {
		defer fmt.Printf("%d ", i)
	}
}
