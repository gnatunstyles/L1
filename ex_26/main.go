package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "qwertyuQiasdfghjkzxcvbnm"
	arr := strings.Split(strings.ToLower(str), "")
	m := make(map[string]bool)
	for i := range arr {
		if !m[arr[i]] {
			m[arr[i]] = true
		} else {
			fmt.Println("false")
			return
		}
	}
	fmt.Println("true")
	return
}
