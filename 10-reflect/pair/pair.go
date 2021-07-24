package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (book *Book) ReadBook() {
	fmt.Println("read a book...")
}

func (book *Book) WriteBook() {
	fmt.Println("write a book...")
}

func main() {
	b := &Book{}
	// 每个对象都存在 一个额 pair，是由多组键值对组成的，其中包含type和value
	// r: pair<type:, value:> 这里未初始化，所以都为空
	var r Reader
	// r: pair<type:Book, value:book{}地址>
	r = b
	r.ReadBook()

	var w Writer
	// r: pair<type:Book, value:book{}地址>
	// 此处断言成功，是因为 type 的类型是 book，是实现了 Writer，w 的实现就是 b = book
	w = r.(Writer)
	w.WriteBook()

}
