package main

import "fmt"

func main() {
	test1()
}
func test1()  {
	//当相关数组还没有定义，可以使用make()函数创建一个切片 同时创建好相关数组
	slice1 := make([]int, 5,10)//type len cap;其中 cap是可选参数
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4
	slice1 = slice1[1:]
	length := len(slice1)
	cap := cap(slice1)//计算切片容量的方法: cap()
	fmt.Printf("length:%d, cap:%d\n", length, cap)
	for i,value := range slice1[1:4]{
		fmt.Printf("Slice at %d is: %d\n", i, value)
	}
}
