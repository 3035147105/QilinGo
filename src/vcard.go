package ghq

import (
	"time"
	"fmt"
)

type Address struct {
	//省
	province string
	//市
	city string
	//区
	district string
}
type VCard struct {
	name string
	address *Address
	birthday time.Time
	//头像
	headPortrait string
}

func TestStruct()  {
	addr := &Address{"山东","淄博","临淄"}
	vcard := new(VCard)
	vcard.name = "王小飞"
	vcard.headPortrait = "图像url"
	vcard.address = addr
	fmt.Println(*vcard.address)
	vcard.birthday = time.Date(2017, time.January, 5, 17,56,10,0, time.UTC)
	fmt.Printf("出生日期：%v,地址：%v\n", vcard.birthday,vcard.address)
}
