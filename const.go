package main

import (
	"fmt"
)

const c0 = 10

const (
	c1 = iota
	c2 = iota
	c3 = iota
)

const c4 = iota

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays //包内私有
)

func main() {
	fmt.Println(c0, c1, c2, c3, c4)
	fmt.Println(numberOfDays)
}
