package main

import "fmt"

//структура
type Television struct{}

//описание интерфейса tv
type tv interface {
	ShowPicture()
}

//Функция структуры Television, принимающая на вход интерфейс tv
func (tv *Television) Show(source tv) {
	source.ShowPicture()
	fmt.Println("Picture is showing")
}

//структура Console
type Console struct{}

//функция PlayGame структуры Console
func (c *Console) PlayGame() {
	fmt.Println("Playing game")
}

//структура

type Antenna struct{}

func (a *Antenna) WatchTV() {
	fmt.Println("Watching TV")
}

//структура SourceConsole
type SourceConsole struct {
	Console
}

//функция ShowPicture структуры SourceConsole

func (sc *SourceConsole) ShowPicture() {
	fmt.Println("Showing picture via playing game")
	sc.PlayGame()

}

//структура SourceAntenna
type SourceAntenna struct {
	Antenna
}

//функция ShowPicture структуры SourceAntenna
func (sa *SourceAntenna) ShowPicture() {
	fmt.Println("Showing picture via Watching TV")
	sa.WatchTV()
}

func main() {
	//создание экземпляров структур SourceAntenna, SourceConsole и Television
	ant := SourceAntenna{}
	con := SourceConsole{}
	tv := Television{}

	//вызов функции Show для демонстрации работы адаптера
	tv.Show(&ant)
	tv.Show(&con)

}
