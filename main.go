package main

import (
	"fmt"
	"study/fucher"
	"study/fucher_postgre/fucher_conected"
	"study/fucker"
)

func main() {
	fmt.Println("Hello, Git")
	fucher.Fucher()
	fucker.Fucker()
	fucher_conected.DataBase()
}
