package main

import (
	"fmt"
	"php4go/string"
)

func main() {
	str := string.StrInstance()

	value := "arash"

	length, err := str.StrLen(value)
	if err != nil {
		fmt.Println("err :", err)
		return
	}

	fmt.Println("length string :", length)

}
