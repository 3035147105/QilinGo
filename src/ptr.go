/**
指针相关，go对指针的支持
 */
package ghq

import (
	"fmt"
)

func TestPtr()  {
	var i1 = 5

	var intP *int//声明一个指针变量，存储一个地址，指向一个位置
	fmt.Printf("指针变量的默认值是：%v\n", intP)

	intP = &i1//“&”取地址符，取出i1的地址赋给intP
	val := *intP//*号放在一个指针变量前面，表示取出该指针指向的地址存储的值

	val = *(&i1)//也可以写成这样

	fmt.Printf("指针变量intP存储的地址:%p，地址对应的值：%d\n", intP, val)
}
//修改数组的值
func Update()  {
	var src = []int{10,11}
	fmt.Printf("修改前的数组：%v\n", src)
	m(&src)
	fmt.Printf("修改后的数组：%v\n", src)
}
//接收一个数组的地址
func m(ptr *[]int)  {
	var array = []int{1,2,3}
	*ptr = array
}
