package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
	sex  byte
}

type Student struct {
	Person  //只有类型 ，没有名字，匿名字段，相当于继承了Person的 ,匿名字段
	name    string
	score   int
	address string
}

func main() {

	var s1 Student = Student{Person{"张三", 20, 0}, "何中", 99, "张家界"}
	//%+v 显示更加详细的
	fmt.Printf("s1 = %+v", s1)

	fmt.Println()
	//指定成员初始化，没有初始化的自动赋值为0
	s2 := Student{Person: Person{name: "李四"}, score: 100} //使用这个方式可以灵活点
	fmt.Printf("s2 = %+v", s2)

	fmt.Println()
	fmt.Println(s2.name)
	var s3 Student
	s3.name = "王五"
	s3.address = "湖头"
	s3.score = 100
	s3.sex = 0
	s3.Person.name = "杜悦"
	fmt.Printf("s3 = %+v", s3)

}
