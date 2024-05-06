package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

// Метод принадлежит Human
func (h *Human) GetName() {
	fmt.Println(h.Name)
}

type Action struct {
	Human				// Встраивание структуры Human
	
	Activity string
}

// Метод принадлежит Action и использует поле из Human
func (a *Action) Which() {
	fmt.Println(a.Name + " " + a.Activity)
}

func main() {
	Vasya := Human{"Vasya", 18}

	step1 := Action{
		Human{"Vasya", 18},
		"wakes up",
	}

	step2 := Action{
		Vasya,
		"stretches",
	}

	step3 := Action{
		Vasya,
		"smiles",
	}

	step1.Which()
	step2.GetName()		// Метод принадлежит Human вызывается от экземпляра Action
	step3.Which()
}
