/**
go 强类型语言，意味着不会进行隐式转换
java 弱类型语言，会隐式转换
 */
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
	fmt.Println("1+2=", f)
	fmt.Println("go允许多个返回值,当实际返回与声明返回不符合时会在运行时报错，编译不报错")
	fmt.Println(ghq.Hello("错误信息", 100))
	fmt.Printf("字符表达式:%s\n", "你好")
	ghq.Time()
	ghq.TestString("gaohaiq","hai")
}
