package main

import "fmt"

//Структура Human
type Human struct {
	Name string
	Hand string
}

//Структура Action со встроенной структурой Human
type Action struct {
	Human
}

//Метод Wave структуры Human
func (h *Human) Wave() {
	fmt.Printf("%s just waved by %s hand!\n", h.Name, h.Hand)
}

func main() {
	//Инициализация переменной action типа Action
	action := Action{
		Human{
			Name: "John",
			Hand: "right",
		},
	}

	//Вызов метода Wave структуры Human внутри переменной action
	action.Wave()

}
