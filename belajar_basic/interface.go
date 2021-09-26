package belajarbasic

import "fmt"

type HashName interface{
	GetName() string
}

type Person struct{
	Nama string
}

func SayHallo(hashname HashName)  {
	fmt.Println("oy", hashname.GetName())
}

func (personal Person) GetName() string {
	return personal.Nama
}

func BelajarInterface()  {
	person := Person{Nama: "jono"}
	SayHallo(person)
}