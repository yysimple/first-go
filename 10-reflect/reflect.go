package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
	// panic: reflect.Value.Interface: cannot return value obtained from unexported field or method
	// weight float64
}

func (t User) Call(user User) {
	fmt.Println("method Call()...")
	fmt.Printf("value =%v\n", t)
}

func (u User) MethodA(user User) {}

func (u User) methodA(user User) {}

func (u User) MethodB() {}

func (u User) methodB() {}

func (u *User) MethodC() {}

func (u User) MethodD(user User) {}

func (u *User) Show(user *User) {
	fmt.Println("user.Id =", user.Id)
	fmt.Println("user.Name =", user.Name)
	fmt.Println("user.Age =", user.Age)
}

func main() {
	user := User{
		Id: 1, Name: "zyy", Age: 18,
	}
	getFieldAndMethod(user)
}

func getFieldAndMethod(all interface{}) {
	// 获取传进来参数的类型
	allType := reflect.TypeOf(all)
	fmt.Println("allType is =", allType.Name())

	// 获取传进来参数的值
	allVaule := reflect.ValueOf(all)
	fmt.Println("allVaule is =", allVaule)

	/**
	allType: 这个是获取到传进来的参数的类型 user
	1. allType.NumField()是长度，代表这个对象有几个字段
	2. field：获取到对应的字段（这里应该是字段的元信息）
	3. value：是字段的 Interface() 获取到的 值
	*/
	for i := 0; i < allType.NumField(); i++ {
		field := allType.Field(i)
		value := allVaule.Field(i).Interface()
		fmt.Println("field =", field, "value =", value)
		fmt.Printf("fieldName = %s, fieldType = %s\n", field.Name, field.Type)
	}

	fmt.Println("---------------------------")

	/**
	1. 获取有多少方法
	2. 获取方法名和类型
	*/
	for i := 0; i < allType.NumMethod(); i++ {
		m := allType.Method(i)
		fmt.Printf("str = %s, value = %v\n", m.Name, m.Type)
	}

	fmt.Println("---------------------------")

	allTypeElem := reflect.TypeOf(&all).Elem()
	for i := 0; i < allTypeElem.NumMethod(); i++ {
		m := allTypeElem.Method(i)
		fmt.Printf("str = %s, value = %v\n", m.Name, m.Type)
	}
}
