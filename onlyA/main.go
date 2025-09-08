package main

import (
	"fmt"

	student "piscine"
)

func main() {
	fmt.Println(student.OnlyA("aaa"))
	fmt.Println(student.OnlyA("aaab"))
	fmt.Println(student.OnlyA(""))
	fmt.Println(student.OnlyA("a"))
}
