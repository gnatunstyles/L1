package main

import "fmt"

func main() {
	arr1 := []string{"a", "f", "d", "s", "c", "x", "z", "v"}
	arr2 := []string{"a", "f", "d", "q", "w", "x", "h", "g"}
	x := []string{}
	exist := make(map[string]bool)
	for _, elt := range arr1 {
		exist[elt] = true
	}
	for _, elt := range arr2 {
		if exist[elt] {
			x = append(x, elt)
		}
	}
	fmt.Println(x)
}
