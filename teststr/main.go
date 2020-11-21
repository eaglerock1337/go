package main

import (
	"fmt"
	"strconv"
)

func main() {
	strVar := "100"
	intVar, _ := strconv.Atoi(strVar)
	fmt.Println(intVar)

	// strVar1 := "-52541"
	// intVar1, _ := strconv.ParseInt(strVar1, 10, 32)

	// strVar2 := "101010101010101010"
	// intVar2, _ := strconv.ParseInt(strVar2, 10, 64)
}
