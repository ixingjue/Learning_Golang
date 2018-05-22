package main

import (
	"time"
	"fmt"
)

var c chan int

func main() {
	c = make(chan int)
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I am waiting,but not too long")
	//time.Sleep(5 * time.Second)
	/*<-c
	<-c*/
	var i int
	L:for{
		select {
		case <-c:
			i++
			if i>1{
				break L
			}

		}
	}
}
func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1
}
