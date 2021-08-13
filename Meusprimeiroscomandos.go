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

	var z int
	z = (2 + 2) * 10
	fmt.Println(z)

}
