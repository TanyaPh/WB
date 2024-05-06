package main

import (
	"fmt"
	"reflect"
)

func matchTypeMethod01(i interface{}) {
	//выбирает case изходя из типа
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan bool:
		fmt.Println("chan bool")
	default:
		fmt.Printf("%T\n", i)
	}
}

func matchTypeMethod02(i interface{}) {
	v := reflect.ValueOf(i)
	fmt.Println(v.Kind())		//получает тип
}

func matchTypeMethod03(i interface{}) {
	fmt.Printf("%T\n", i)		//выводит тип
}

func main() {
	sl := []interface{}{4, "go", false, make(chan int), make(chan bool), struct{}{}}
	for _, v := range sl {
		matchTypeMethod01(v)
		matchTypeMethod02(v)
		matchTypeMethod03(v)
	}
}
