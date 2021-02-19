package main

import (
	"fmt"
	"oop"
)

func main() {

	var s1 oop.Student = oop.Student{}
	s1.SetName("张三")
	//%+v 显示更加详细的
	name := s1.GetName()
	fmt.Println(name)
	//匿名结构体
	oop.Amonymous()
}
