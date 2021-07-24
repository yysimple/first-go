package main

import "fmt"

// 声明一种数据类型，然后用别名替代
type myInt int

// 定义一个结构体 Book
type Book struct {
	title    string
	pageSize int
}

// 改变对象的值
func updateBookFalse(book Book) {
	// 这里的book只是副本，所以只有方法里面的值是修改了的
	book.pageSize = 1
	fmt.Printf("updateBookFalse.book = %v\n", book)
}

// 传递的是对象的引用地址
func updateBookTrue(book *Book) {
	book.title = "兄弟"
	fmt.Printf("updateBookTrue.book = %v\n", book)
}

func main() {
	var age myInt = 10
	// type = main.myInt, 这里打印出来的就是我们自己定义的数据类型的别名
	fmt.Printf("type = %T\n", age)

	var book Book
	book.title = "平凡的世界"
	book.pageSize = 1000

	fmt.Printf("book = %v\n", book)

	fmt.Println("-------------")

	updateBookFalse(book)
	// 这里的值是不会变化的
	fmt.Printf("book = %v\n", book)

	fmt.Println("-------------")

	updateBookTrue(&book)
	// 这里的值是会变化的，和方法里面打印的结果是一样的，方法里面会多个 &
	fmt.Printf("book = %v\n", book)

}
