/**
 */
package main

import (
	"fmt"
	//"strings"
)

var a = "G"
var b string
var num1 int
var e error

func main() {
	//var str string = "This is an example of a string"
	//fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	//fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	//n()
	//m()
	//f()
	//num1, e = add(1,2)
	//fmt.Printf("自定义返回值%d和错误%p\n", num1, e)
	//n := 0
	//reply := &n//&取地址符，返回相应变量的内存地址
	//Multiply(3,4, reply)
	//fmt.Println(reply)//0xc82000a310：内存地址表示方法
	//fmt.Println(*reply)//12
	//testFor()
	leijia()
}
//使用:=可以高效的创建一个新的变量，局限性是只能在函数中使用
func bingxingfuzhi()  {
	c, d, e := 5, 7, "abc"//优雅的并行赋值
	fmt.Print(c,d,e)
}
func n()  {
	print("\n")
	print(a)
}
func m()  {
	a = "O"
	print(a)
}
func f()  {
	print("\n")
	b = "G"
	print(b)
	f1()
}
func f1()  {
	b := "O"
	print(b)
	f2()
}
func f2()  {
	print(b)
}

func add(a ,b int)(c int, e error){
	c = a+b
	return c , e
}
//全局变量操作
func Multiply(a,b int, reply *int) {
	*reply = a*b
}
//for-range结构
//char是索引位置的拷贝，不能用来修改该索引位置的值
func testFor()  {
	array := []string{"ami", "tom", "jack"}
	for pos,char := range array{
		fmt.Printf("Character on position %d is: %s\n", pos, char)
	}
}
//改变数组的值
func leijia()  {
	items := []int{10,20,30,40,50}
	//将数组的值都变为原来的两倍
	//这段代码起不到预定的作用
	for _,item := range items{
		item *= 2
	}
	items = getSlice(items)
	fmt.Printf("items为：%v\n",items)
	items = append(items, 1,2,3,4,5)
	fmt.Printf("追加后，items为：%v\n", items)
	slicecp := make([]int, len(items))
	copy(slicecp, items)
	fmt.Printf("复制后，新的slicecp为：%v", slicecp)
}
//接受一个切片，返回一个切片
func getSlice(items []int) (res []int){
	//这段代码起了作用
	for index := range items{
		items[index] = items[index] * 2
	}
	res = items
	return
}


