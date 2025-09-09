package main

import (
	"fmt"

	"github.com/DmitriyVTitov/pointer"
)

type T struct {
	BoolPointer   *bool
	StringPointer *string
}

func main() {
	var t = T{
		BoolPointer:   pointer.To(true),
		StringPointer: pointer.To("string"),
	}
	_ = t

	var a any
	var nilPointer = pointer.To(a)
	_ = nilPointer

	s := "ABC"
	val := pointer.Val(&s)
	fmt.Println(val == s) // true

	var empty *string
	fmt.Println(pointer.Val(empty) == "") // true
}
