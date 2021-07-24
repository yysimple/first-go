package main

import (
	"fmt"
	"reflect"
)

// `` 号里面的数据就是 字段的tag内容，这里注意的是 info: 这个后面不能有空格
type Book struct {
	name     string `info:"书的名字" desc:"余华的书"`
	pageSize int    `info:"书纸张的页数"`
}

func getTags(all interface{}) {
	allType := reflect.TypeOf(all)
	for i := 0; i < allType.NumField(); i++ {
		field := allType.Field(i)
		tag := field.Tag
		info := tag.Get("info")
		desc := tag.Get("desc")
		fmt.Println("info: ", info, "desc: ", desc)
	}
}

func getElemTags(all interface{}) {
	// Elem returns a type's element type，返回元素的类型，user对应的就是user
	allType := reflect.TypeOf(all).Elem()
	for i := 0; i < allType.NumField(); i++ {
		field := allType.Field(i)
		tag := field.Tag
		info := tag.Get("info")
		desc := tag.Get("desc")
		fmt.Println("info: ", info, "desc: ", desc)
	}
}

func main() {
	// var book = Book{name: "zs", pageSize: 123}
	var book Book
	getTags(book)
	// reflect.(*rtype).Elem(0xbd88e0) 这种写法需要传入执行类型
	getElemTags(&book)
}
