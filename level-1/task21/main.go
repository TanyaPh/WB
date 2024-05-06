package main

import "fmt"

type racingCar interface {
	rushing()
}

type racingBolide struct { }

// Реализация интерфейса
func (r *racingBolide) rushing() {
	fmt.Println("I'm going very fast.")
}

type car struct { }

// Не реализует интерфейс
func (c *car) going() {
	fmt.Println("I'm going.")
}

func race(car racingCar) {
	car.rushing()
}

type adapter struct{
    *car
}

// реализуем метод интерфейса
func (a *adapter) rushing() {
	a.going()
}

func newAdapter(c *car) racingCar {
    return &adapter{c}
}

func main() {
	bolide := &racingBolide{}
	car := &car{} 

	race(bolide)
	race(newAdapter(car))
}
