package main

import (
	"fmt"
	"./src"
)
func main() {
	var str string
	var res int
	var f float64
	str = ghq.ReturnStr()
	fmt.Println("输出一句话")
	fmt.Println(str)
	fmt.Println("作一个加法运算")
	res = ghq.Add(1,2)
	fmt.Println(res)
	fmt.Println("作一个减法运算");
	sub := ghq.Sub(10,20)
	fmt.Println("1-2: %d", sub);
	f = 1 +2;
	fmt.Println(f)
}
