package ghq

import (
	"math"
	"time"
	"fmt"
	"strings"
)
func Sub(a float64,b float64) float64 {
	max := a - b
	res := math.Abs(max)
	return res
}
//go中 时间的方法
func Time()  {
	t := time.Now()
	day := t.Day()
	month := t.Month()
	year:= t.Year()
	hour:=t.Hour()
	minute:=t.Minute()
	second:=t.Second()
	fmt.Printf("%d年%d月%d日 %d时%d分%d秒\n", year,month,day,hour,minute,second)
}
//go中 操作字符串的包
func  TestString(s1 string, s2 string){
	fmt.Printf("%s中是否包含%s吗? %v\n",s1,s2,strings.Contains(s1, s2))
}

//循环一个list

