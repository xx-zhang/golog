package tests

import (
	"fmt"
	"reflect"
)


func TestStruct(a interface{})  {

	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a) //获取reflect.Type类型

	kd := val.Kind() //获取到a对应的类别
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取到该结构体有几个字段
	num := val.NumField()
	fmt.Printf("该结构体有%d个字段\n", num) //4个

	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d:值=%v\n",i,val.Field(i))
		//获取到struct标签，需要通过reflect.Type来获取tag标签的值
		keyVal := typ.Field(i).Tag.Get("json")
		//如果该字段有tag标签就显示，否则就不显示
		if keyVal != ""{
			fmt.Printf("Field %d:key=%v\n", i, keyVal)
		}
	}
}
