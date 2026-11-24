package main

import (
	"fmt"
)

func add(x int,y int ) int {
	return x + y
}

func add1(x,y int) int {
	return x + y
}

func main() {
	fmt.Println(add(2,3))
	fmt.Println(add1(2,4))
}