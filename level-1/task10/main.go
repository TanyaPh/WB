package main

import "fmt"

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := map[int][]float64{}

	for _, v := range data {
		k := int(v/10) * 10			//вычесление группы с шагом в 10 
		m[k] = append(m[k], v)		//добавление к группе
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
