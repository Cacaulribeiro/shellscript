package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	x := 23
	fmt.Println("Esse é o número ", x)

	var y int
	y = 2 + 2 + 3
	fmt.Println(y)

}
