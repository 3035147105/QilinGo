/**
 */
package main

import (
	"fmt"
	"strings"
)

var a = "G"
var b string

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}
//使用:=可以高效的创建一个新的变量，局限性是只能在函数中使用
func bingxingfuzhi()  {
	c, d, e := 5, 7, "abc"//优雅的并行赋值
	fmt.Print(c,d,e)
}
func n()  {
	print(a)
}
func m()  {
	a = "O"
	print(a)
}
func f()  {
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


