package main

import (
	"fmt"
	"mymodule/mypackage"
)

func main() {
	hello := mypackage.SayHello()
	fmt.Println(hello)
}
