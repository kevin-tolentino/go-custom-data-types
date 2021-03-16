package main

import (
	"go-custom-data-types/organization"
)

func main() {
	p := organization.NewPerson("Kevin", "Tolentino")
	println(p.ID())
	println(p.FullName())
	//println(p.FullName())
}