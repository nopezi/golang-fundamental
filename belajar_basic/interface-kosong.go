package belajarbasic

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return "True"
	} else {
		return "Ops"
	}
}

func InterfaceKosong()  {
	var data interface{} = Ups(3)
	fmt.Println("interface kosong", data)
}