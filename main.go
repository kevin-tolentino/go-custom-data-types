package main

import (
	"fmt"
	"go-custom-data-types/organization"
)

func main() {
	p := organization.NewPerson("Kevin", "Tolentino")
	err := p.SetTwitterHandler("@kev_tol")
	if err != nil {
		fmt.Printf("an error occured setting twitter handler: %s\n", err.Error())
	}
	println(p.TwitterHAndler())
	println(p.ID())
	println(p.FullName())
	println(p.FullName())
}
