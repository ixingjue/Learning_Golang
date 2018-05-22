package main

import "fmt"

//栈
type stack struct {
	i    int
	data [10] int
}

//todo 第三章包的学习
func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++

}
func (s *stack) pop(k int) {
	s.i--
	return s.data[s.i]

}
func main() {
	var s stack
	s.push(25)
	s.push(14)
	fmt.Printf("stack %v\n", s)
}
