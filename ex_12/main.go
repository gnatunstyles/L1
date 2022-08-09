package main

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)
	for _, elt := range arr {
		set[elt] = true
	}
	for key := range set {
		println(key)
	}
}
