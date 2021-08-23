package main

import (
	"fmt"
)

func Calculate(x int) int {
	return x+2
}

func Add(x,y int) int {
	return x+y
}

func main()  {
	fmt.Println("testdata")
	res:=Calculate(2)
	fmt.Println(res)
}