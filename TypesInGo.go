package main

import "fmt"

func isSameSt(a, b string) bool {
	fmt.Println(a, b)
	return a == b
}

func isSameInt(a, b int) bool {
	fmt.Println(a, b)
	return a == b
}

func main() {
	var a, b string = "hello", "hello"
	fmt.Println("are A and B same? ", isSameSt(a, b))
	var c, d int = 14, 55
	fmt.Println("are A and B same? ", isSameInt(c, d))
	p := &a
	fmt.Println(*p)
	*p = "world"
	fmt.Println(len(a), a, len(a)*len(*p))
	fmt.Println(p)
}
