package main

import (
	"fmt"
	"oop"
)

func main() {

	var s1 = oop.Student{}
	s1.SetName("张三")
	//%+v 显示更加详细的
	s1.GetName()

	//指定成员初始化，没有初始化的自动赋值为0
	s2 := oop.Student{} //使用这个方式可以灵活点
	fmt.Printf("s2 = %+v", s2)

}
