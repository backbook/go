package main

import "fmt"

func main() {

	i := 10

	anonymousF1 := func() {
		fmt.Println(i)
	}
	anonymousF1()

	type funcType func()

	var f2 funcType

	f2 = anonymousF1

	f2()

	func(a int) {
		fmt.Printf("this is %c", a)
	}(99)
	fmt.Println()

	//匿名函数有参数返回值
	max, min := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 23)

	fmt.Println(max, min)

}
