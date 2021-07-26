package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (t *Human) Eat() {
	fmt.Println("Human.Eat() ...")
}

func (t *Human) Walk() {
	fmt.Println("Human.Walk() ...")
}

// 子类
type SuperMan struct {
	// 这样写就继承了父类
	Human
	level int
}

func (t *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat() ...")
}

func (t *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly() ...")
}

func main() {
	human := Human{name: "wcx", age: 18}

	human.Eat()
	human.Walk()

	fmt.Println("-------------------")
	superMan := SuperMan{Human{name: "zyy", age: 18}, 1}

	superMan.Walk() // 这里是父类的方法
	superMan.Eat()  // 这里是 重写父类的方法
	superMan.Fly()  // 这里是子类的方法
}
