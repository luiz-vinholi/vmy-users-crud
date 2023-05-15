package main

import (
	"github.com/luiz-vinholi/vmy-users-crud/src"
	"github.com/luiz-vinholi/vmy-users-crud/src/interfaces/rest"
)

func main() {
	Run()
}

func Run() {
	src.Init()
	rest.Run()
}
