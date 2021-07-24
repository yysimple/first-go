package main

import "fmt"

// 这里注意的是，类名大写的话，其他的包就能访问当前类
type Hero struct {
	// 属性也是一样的，只要是大写其他包据可以访问到，这里没有访问权限修饰符的说法
	Name  string
	Age   int
	Level int
}

// 可以理解为，这里加了Hero，就是声明这是Hero 对象的方法,get方法就是当前对象的值，所以使不使用指针都是一样的
func (t *Hero) GetName() string {
	return t.Name
}

// 设置值，不加 * ，同样只会作用在方法内
// 这里就是 java 中的 set方法，在java中，传对象的时候，都是传递的是内存地址，
// 而go中则是副本，是不会去修改原有的值的，所以需要加上 *，去玩成同样的功能
func (t *Hero) SetName(Name string) {
	t.Name = Name
}

func (t *Hero) Show() {
	fmt.Println("Name =", t.Name)
	fmt.Println("Age =", t.Age)
	fmt.Println("Level =", t.Level)
}

func main() {
	heroZs := Hero{Name: "zs", Age: 18, Level: 1}
	// 打印对象
	heroZs.Show()
	fmt.Println("GetName.Name =", heroZs.GetName())
	heroZs.SetName("张三")
	fmt.Println("SetName.Name =", heroZs.GetName())

	fmt.Println("-------------------------")

	heroLs := Hero{Name: "Ls", Age: 20, Level: 2}
	heroLs.Show()
}
