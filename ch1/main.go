package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	x = 10
}
