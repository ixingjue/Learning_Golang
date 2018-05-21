package main

import "fmt"

func main() {
	//编写一个程序，打印从 1 到 100 的数字。当是三个়数就打印 “Fizz” 代替数字，当是ࡃ的়数就打印 “Buzz”。当数字同时是三和ࡃ的়数 时，打印 “FizzBuzz”。
	const (
		FIZZ = 3
		BIZZ = 5
	)
	var p bool
	for i := 1; i < 100; i++ {
		p = false
		if i%FIZZ == 0 {
			fmt.Printf("Fizz")
			p = true
		}
		if i%BIZZ == 0 {
			fmt.Printf("Bizz")
			p = true
			//如果能被 BUZZ 整除，打印 “Buzz”。注意，FizzBuzz 的情况已经被处理了；
		}
		if !p {
			fmt.Printf("%v", i)
		}
		fmt.Println()
	}
}
