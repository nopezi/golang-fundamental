package belajarbasic

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"nama": name,
		}
	}
}

func Nill()  {
	var orang map[string]string = newMap("jono")
	if orang == nil {
		fmt.Println("data orang kosong")
	} else {
		fmt.Println("data nill", orang)
	}
}