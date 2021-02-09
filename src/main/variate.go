package main

import (
	"fmt"
	"math"
)

var m int = 100

func varTest() {
	//批量定义用var()
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)

	//var 变量名 变量类型 = 10
	var e int = 10
	var f string = "123"
	var _a float64 = 131.2
	fmt.Println(e, f, _a)

	var age, name, height, debt = 19, "张三", "1314", 10.231
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(height)
	fmt.Println(debt)
	//推荐使用在debug,官网推荐使用fmt.Println()
	print(age)

	// 短变量声明
	//只能在函数内部使用，并非全局的函数编译会失败 := 是省略了 var
	number := 100
	fmt.Println(number)
	//比较m的定义
	fmt.Println(m)

	const (
		n1 = iota //0
		n2 = iota //1
		_         //2
		n3        //3
		n4        //4
	)

	//计算换算小规律
	const (
		B  = 1 << (10 * iota) // 0
		KB = 1 << (10 * iota) //1
		MB = 1 << (10 * iota) //2
		GB = 1 << (10 * iota) //3
		TB = 1 << (10 * iota) //4
		PB = 1 << (10 * iota) //4
	)
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)

	//整型描述
	//在使用int类型时候，int64是java中long，int16是short int8是byte
	var number1 int64 = 100
	fmt.Println(number1)
	fmt.Println(math.Pow(2, 31) - 1)

	var v = 0b00101010
	var s = 123_123
	var t = 0o131
	fmt.Println(v)
	fmt.Println(t)
	fmt.Println(s)

	var str = `
     hello
     world
     世界`
	fmt.Println(str)

	var t1 int = '中'
	fmt.Println(t1)

	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
}

func main() {

	//varTest()
	traversalString()

}

func traversalString() {

	s := "hello沙河"
	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Println(r)
	}

}
