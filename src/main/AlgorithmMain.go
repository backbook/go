package main

import (
	"algorithm"
)

func main() {
	// :=  局部变量，类型自动推导
	recursive := algorithm.Recursive(10)
	println(recursive)
}
