package obj

type Person struct {
	name string
	age  int
	sex  byte
}

type Student struct {
	Person  //只有类型，没有名字，匿名字段，相当于继承了Person的 ,匿名字段
	score   int
	address string
}
