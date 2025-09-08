package main

import (
	"fmt"

	student "piscine"
)

func main() {
	fmt.Println(student.OnlyZ("zzz"))
	fmt.Println(student.OnlyZ("zzzb"))
	fmt.Println(student.OnlyZ(""))
	fmt.Println(student.OnlyZ("z"))
}