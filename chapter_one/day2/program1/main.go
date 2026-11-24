package main

import (
	"fmt" 
	"math/rand"
)

func main(){
	 pesudoRN()
	 deferTest()
}

func pesudoRN() {
	fmt.Println("my favorite number is :",rand.Intn(10));
}

func deferTest(){
	fmt.Println("counting")
	for i:=0; i < 10;i++{
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
