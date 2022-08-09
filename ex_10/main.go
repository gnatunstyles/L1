package main

import (
	"fmt"
	"math"
)

func main() {
	var temps = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // массив с тестовымси значениями
	groups := make(map[int][]float64)                                        //структура для хравнения распределенных данных
	for _, elt := range temps {                                              // итерация по массиву тестовых значений
		fmtElt := int(math.Trunc(elt)) / 10 * 10               //определение подгруппы текущего элемента
		groups[int(fmtElt)] = append(groups[int(fmtElt)], elt) //запись элемента в map по ключу подгруппы
	}
	fmt.Println(groups)
}
