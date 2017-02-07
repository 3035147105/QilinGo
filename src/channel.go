/**
信道相关，go在语法层面支持并发
 */
package ghq

import (
	"fmt"
)

func TestChan()  {
	ch := make(chan string) //初始化一个存储string类型的信道
	go func(s string) { //定义一个匿名函数并用协程的方式跑
		ch <- s  //往信道里添加数据
	}("Ping!") //调用匿名函数

	fmt.Println(<-ch) //取数据
}

func loop(){
	for i:=0; i<10; i++ {
		fmt.Printf("%d", i)
	}
}
//-------------------------------------分割线------------------------------
var count = 10
var chh chan int= make(chan int,count)//初始化一个存储int类型，容量为10的信道
func foo1(id int)  {
	fmt.Printf("子线程%d的方法执行完了\n", id)

	//向chh信道中加数据，如果没有其他goroutine来取走这个数据，那么挂起foo1,直到有goroutine把这个数据取走
	chh <- id//上边这行注释非常重要，信道使用的关键点
	//逻辑代码应该放在往chh里放数据之前
}
func TestChan1()  {
	fmt.Println("进入主线程")
	for i:=0; i<count; i++{
		go foo1(i) //这一句会去执行fool方法，但是fool方法的通道是会阻塞住的，如果
	}
	for i:=0; i<count; i++{
		//从chh中取数据，如果chh中还没放数据，那就挂起当前线程，直到有goroutine放数据为止
		fmt.Println(<- chh)
	}
	fmt.Println("主线程继续执行")
}

