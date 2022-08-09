package main

import "fmt"

type Television struct{}

type television interface {
	ShowPicture()
}

func (tv *Television) Show(source television) {
	source.ShowPicture()
	fmt.Println("Picture is showing")
}

type Console struct{}

func (c *Console) PlayGame() {
	fmt.Println("Playing game")
}

type Antenna struct{}

func (a *Antenna) WatchTV() {
	fmt.Println("Watching TV")
}

type SourceConsole struct {
	Console
}

func (sc *SourceConsole) ShowPicture() {
	fmt.Println("Showing picture via playing game")
	sc.PlayGame()

}

type SourceAntenna struct {
	Antenna
}

func (sa *SourceAntenna) ShowPicture() {
	fmt.Println("Showing picture via Watching TV")
	sa.WatchTV()
}

func main() {
	ant := SourceAntenna{}
	con := SourceConsole{}

	tv := Television{}
	tv.Show(&ant)
	tv.Show(&con)

}
