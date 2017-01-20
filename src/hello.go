package ghq

func ReturnStr() string {
	return "Hello go!"
}

func Add(a int,b int) int{
	return a+b
}
//go允许有多个返回值
func Hello(param1 string, param2 int)(error string, ret1 int, ret2 int)  {
	return param1, param2, 0;
}
