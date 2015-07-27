package main

import (
	"fmt"
)

type person struct {
	Id      string
	Name    string
	Address string
}

func main() {
	var personDb map[string]person
	personDb = make(map[string]person, 100)
	personDb["123"] = person{"1", "fanzw", "xxxxx"}
	//or yMap = map[string] PersonInfo{
	//"1234": PersonInfo{"1", "Jack", "Room 101"},
	//}
	fmt.Println(personDb)
	value, ok := personDb["123"]
	if ok { // 找到了
		fmt.Println(value) // 处理找到的value
	}
	delete(personDb, "123")
	fmt.Println(personDb)
}
