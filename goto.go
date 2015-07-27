package main

import (
	"fmt"
)

func main() {
	i := 0
GOTO:
	fmt.Println(i)
	i++
	if i < 10 {
		goto GOTO
	}
}
