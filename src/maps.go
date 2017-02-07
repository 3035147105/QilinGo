package ghq

import "fmt"

type Person struct {
	Name string
	Address string
}
func TestMap()  {
	myMap := make(map[string] Person, 2)
	//增加数据
	myMap["1234"] = Person{"Tom", "America"}
	myMap["123"] = Person{"Jack", "Japanese"}
	//查询
	person, ok := myMap["123"]
	if ok{
		fmt.Printf("找到了Id为123的person:%v\n", person)
	}else {
		fmt.Println("Id为123的person不存在")
	}
	//修改
	myMap["123"] = Person{"ghq", "China"}
	//person.Name = "ghq"//这样修改不会起作用
	//myMap["123"].Name = "ghq"//编译时会报错
	//遍历map
	list(myMap)
	//删除map,使用go内置的delete方法
	delete(myMap, "1234")
	list(myMap)

}
//遍历map
func list(myMap map[string] Person)  {
	for i,v := range myMap{
		fmt.Printf("index为%s的person为：%v\n", i, v)
	}
}