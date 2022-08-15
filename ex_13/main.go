package main

func main() {
	a := 1
	b := 2
	println(a, b)
	// меняем местами значения a и b
	a, b = b, a
	println(a, b)
}
