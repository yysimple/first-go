package main

import "fmt"

type Animal interface {
	Sleep()
	Color() string
	Type() string
}

type Cat struct {
	color string
}

func (cat *Cat) Sleep() {
	fmt.Println("Cat.Sleep()...")
}

func (cat *Cat) Color() string {
	fmt.Println("Cat.Color()...", cat.color)
	return cat.color
}

func (cat *Cat) Type() string {
	fmt.Println("Cat.Type()...")
	return "cat"
}

type Dog struct {
	color string
}

func (dog *Dog) Sleep() {
	fmt.Println("Dog.Sleep()...")
}

func (dog *Dog) Color() string {
	fmt.Println("Dog.Color()...", dog.color)
	return dog.color
}

func (dog *Dog) Type() string {
	fmt.Println("Dog.Type()...")
	return "dog"
}

func ShowColor(animal Animal) {
	animal.Sleep()
	fmt.Println("Animal.color =", animal.Color())
}

func main() {
	var animal Animal
	animal = &Cat{color: "white"}
	animal.Sleep()

	fmt.Println("--------------")

	animal = &Dog{color: "Red"}
	animal.Sleep()

	fmt.Println("--------------")

	// 这里就是多态的思想，这里需要加 &，是因为上面的实现需要传递的是指针
	ShowColor(&Cat{color: "blue"})
	ShowColor(&Dog{color: "yellow"})

}
