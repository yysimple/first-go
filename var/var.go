package main

import "fmt"

var ageA = 18
var nameB string = "global"

// 这种方式不允许定义全局变量
// ageC := 10

func main() {
	// 1. 声明一个变量，默认值为0
	var ageZs int
	fmt.Println("ageZs = ", ageZs)
	fmt.Printf("类型为：%T\n", ageZs)

	var nameZs string
	fmt.Println("nameZs = ", nameZs)
	fmt.Printf("类型为：%T\n", nameZs)

	fmt.Println("---------------------------")

	// 2. 声明一个变量，自己指定初始化值
	var ageLs int = 18
	fmt.Println("ageLs = ", ageLs)
	fmt.Printf("类型为：%T\n", ageZs)

	var nameLs string = "李四"
	fmt.Println("nameLs = ", nameLs)
	fmt.Printf("类型为：%T\n", nameLs)

	fmt.Println("---------------------------")

	// 3. 声明一个变量，在初始化的时候可以省去类型关键字，通过值会自动匹配
	var ageWw = 19
	fmt.Println("ageWw = ", ageWw)
	fmt.Printf("类型为：%T\n", ageWw)

	var nameWw = "王五"
	fmt.Println("nameWw = ", nameWw)
	fmt.Printf("类型为：%T\n", nameWw)

	fmt.Println("---------------------------")

	// 4. 声明一个变量，使用 ： ，这样可以省略 var，自动推断出其类型
	ageZl := 20
	fmt.Println("ageZl = ", ageZl)
	fmt.Printf("类型为：%T\n", ageZl)

	nameZl := "赵六"
	fmt.Println("nameZl = ", nameZl)
	fmt.Printf("类型为：%T\n", nameZl)

	gradeZl := 99.99
	fmt.Println("gradeZl = ", gradeZl)
	fmt.Printf("类型为：%T\n", gradeZl)

	fmt.Println("---------------------------")

	// 5. 全局变量的使用方式
	fmt.Println("ageA = ", ageA, "nameB = ", nameB)
	fmt.Printf("ageA的值为：%s, nameB的值为：%s\n", ageA, nameB)
	fmt.Printf("ageA类型为：%T, nameB类型为：%T\n", ageA, nameB)

	fmt.Println("---------------------------")

	// 6. 声明一个变量，多行声明变量的方式
	var ageSy, nameSy = 18, "孙一"
	fmt.Println("ageSy = ", ageSy, "nameSy = ", nameSy)

	var (
		ageQe  int    = 15
		nameQe string = "前二"
	)
	println("ageQe = ", ageQe, "nameQe = ", nameQe)

}
