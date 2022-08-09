package main

var justString string

func createHugeString(n int) string {
	return "bigstringsdfsdfaassasassdfasdfasdfasdfasdfasdfasdfasfasdfasfasdfasfasfasfasdfasdfasdfasdfasdfasdfasdf"
}

func someFunc() {
	v := createHugeString(1 << 10)
	//плохой вариант, так как строка может некорректно выделиться при выборе строки с символами utf-8
	// justString = v[:100]
	justString = string([]rune(v)[:100])
}

func main() {
	someFunc()
}
