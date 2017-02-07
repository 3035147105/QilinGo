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
var ch = make(chan int)

func fo(id int)  {
	fmt.Printf("子线程%d的方法开始执行\n", id)
	ch <- id
	fmt.Printf("子线程%d的方法执行完了\n", id)
}
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
	//leijia()
	/**
	    在主线程中只往信道里塞数据，而不取数据会报死锁错误
	    写到另一个方法fo中，然后用go fo()的方式去调用就不会出错
	    原因是否因为 信道会阻塞当前线程，阻塞的是fo，对main方法没有影响？
	 */
	fmt.Println("this line code is run?")
	//go fo(1)
	//go fo(2)
	//go fo(3)
	ch  <- 1
	fmt.Println("this line code wont run")
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


