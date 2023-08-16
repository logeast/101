package main

import (
	"fmt"
	"mypackage/greet"
)

func main() {
	hello := greet.SayHello()
	fmt.Println(hello)
}
