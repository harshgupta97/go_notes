package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "hello",
	}

	fmt.Printf("co = {num: %v, str :%v}\n", co.num, co.str)
	fmt.Printf("also num = %v \n", co.base.num)

	fmt.Printf("describe = %v \n", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
