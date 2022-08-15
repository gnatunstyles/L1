package main

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)
	//создания карты типа с ключом string и значением bool
	for _, elt := range arr {
		set[elt] = true
	}
	//вывод всех ключей созданной карты
	for key := range set {
		println(key)
	}
}
