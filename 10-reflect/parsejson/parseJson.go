package main

import (
	"encoding/json"
	"fmt"
)

type Fruit struct {
	Name  string  `json:"name"`
	Price float64 `json:"rmb"`
	Sweet int     `json:"level"`
}

func main() {
	apple := Fruit{Name: "apple", Price: 10.8, Sweet: 1}
	// 这里是转换成json字符串
	appleJson, err := json.Marshal(apple)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Printf("appleJson = %s\n", appleJson)

	fmt.Println("---------------------------")

	newApple := Fruit{}
	// err = json: Unmarshal(non-pointer main.Fruit), 这里是需要传入指针的
	err = json.Unmarshal(appleJson, &newApple)
	if err != nil {
		fmt.Println("err =", err)
		return
	}

	fmt.Printf("value = %v\n", newApple)

}
