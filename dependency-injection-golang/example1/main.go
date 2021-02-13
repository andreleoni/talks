package main

import "fmt"

type Store interface {
	Sell() bool
}

type EnjuPro struct{}

func (EnjuPro) Sell() bool {
	return true
}

func main() {
	var currentStore Store

	currentStore = EnjuPro{}

	if currentStore.Sell() {
		fmt.Println("Item vendido")
	} else {
		fmt.Println("Não vendido")
	}
}
