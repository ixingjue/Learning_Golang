package main

import "fmt"

func main() {
	rec(5)
}
func rec(i int) {
	if i == 10 {
		return
	}
	rec(i + 1)
	fmt.Printf("%d\n", i)

}
